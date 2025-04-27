package main

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

type Client struct {
	Address  *net.UDPAddr
	LastSeen time.Time
	Active   bool
	mu       sync.Mutex
}

func (c *Client) setActive(active bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Active = active
	if active {
		c.LastSeen = time.Now()
	}
}

func (c *Client) isActive() bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.Active
}

type Server struct {
	conn        *net.UDPConn
	clients     map[string]*Client
	clientsLock sync.RWMutex
}

func NewServer(address string) (*Server, error) {
	addr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		return nil, err
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		return nil, err
	}

	return &Server{
		conn:    conn,
		clients: make(map[string]*Client),
	}, nil
}

func (s *Server) RegisterClient(address string) error {
	udpAddr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		return err
	}

	s.clientsLock.Lock()
	s.clients[address] = &Client{
		Address: udpAddr,
		Active:  false,
	}
	s.clientsLock.Unlock()

	return nil
}

func (s *Server) Start() {
	// Start goroutine to listen for client responses
	go s.listenForResponses()

	// Start goroutine to periodically ping clients
	go s.pingClientsRoutine()
}

func (s *Server) listenForResponses() {
	buffer := make([]byte, 1024)
	for {
		n, addr, err := s.conn.ReadFromUDP(buffer)
		if err != nil {
			log.Printf("Error reading from UDP: %v", err)
			continue
		}

		// Process the message
		message := string(buffer[:n])
		clientKey := addr.String()

		if message == "ECHO-RESPONSE" {
			s.clientsLock.RLock()
			client, exists := s.clients[clientKey]
			s.clientsLock.RUnlock()

			if exists {
				client.setActive(true)
				log.Printf("Client %s marked as active", clientKey)
			} else {
				// New client responded, let's add it
				s.clientsLock.Lock()
				s.clients[clientKey] = &Client{
					Address:  addr,
					LastSeen: time.Now(),
					Active:   true,
				}
				s.clientsLock.Unlock()
				log.Printf("New client %s registered and marked as active", clientKey)
			}
		}
	}
}

func (s *Server) pingClientsRoutine() {
	for {
		s.pingAllClients()
		time.Sleep(30 * time.Second)
	}
}

func (s *Server) pingAllClients() {
	s.clientsLock.RLock()
	clientList := make([]*Client, 0, len(s.clients))
	for _, client := range s.clients {
		clientList = append(clientList, client)
	}
	s.clientsLock.RUnlock()

	for _, client := range clientList {
		go s.pingClient(client)
	}
}

func (s *Server) pingClient(client *Client) {
	message := []byte("ECHO-REQUEST")
	received := make(chan bool, 1)

	// Try up to 3 times with 3-second timeout
	for attempt := 1; attempt <= 3; attempt++ {
		// Send echo request
		_, err := s.conn.WriteToUDP(message, client.Address)
		if err != nil {
			log.Printf("Failed to send echo to %s: %v", client.Address, err)
			client.setActive(false)
			return
		}

		log.Printf("Sent echo request to %s (attempt %d)", client.Address, attempt)

		// Wait for response or timeout
		select {
		case <-received:
			// Response received in listenForResponses(), we're done
			return
		case <-time.After(3 * time.Second):
			if attempt == 3 {
				// Third attempt failed, mark client as inactive
				client.setActive(false)
				log.Printf("Client %s marked as inactive after 3 attempts", client.Address)
			} else {
				// Retry
				log.Printf("Retrying echo to %s (attempt %d failed)", client.Address, attempt)
			}
		}
	}
}

func (s *Server) PrintClientStatus() {
	s.clientsLock.RLock()
	defer s.clientsLock.RUnlock()

	fmt.Println("Client Status:")
	for addr, client := range s.clients {
		status := "inactive"
		if client.isActive() {
			status = "active"
		}
		fmt.Printf("%s: %s (Last seen: %s)\n", addr, status, client.LastSeen)
	}
}

func main() {
	server, err := NewServer(":8053")
	if err != nil {
		log.Fatalf("Failed to create server: %v", err)
	}

	// Register some clients (in a real application, clients might register themselves)
	// These are example clients you would replace with actual client addresses
	server.RegisterClient("127.0.0.1:8054")
	server.RegisterClient("127.0.0.1:8055")

	// Start the server
	server.Start()

	// Print status every minute
	ticker := time.NewTicker(1 * time.Minute)
	for range ticker.C {
		server.PrintClientStatus()
	}
}
