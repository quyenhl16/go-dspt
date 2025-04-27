package main

import (
	"flag"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	port := flag.String("port", "8054", "UDP port to listen on")
	respond := flag.Bool("respond", true, "Whether to respond to echo requests")
	flag.Parse()

	addr, err := net.ResolveUDPAddr("udp", ":"+*port)
	if err != nil {
		log.Fatalf("Failed to resolve UDP address: %v", err)
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	defer conn.Close()

	log.Printf("Client listening on port %s, will respond: %v", *port, *respond)

	// Handle graceful shutdown
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigCh
		log.Println("Shutting down...")
		conn.Close()
		os.Exit(0)
	}()

	buffer := make([]byte, 1024)
	for {
		n, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			log.Printf("Error reading from UDP: %v", err)
			continue
		}

		message := string(buffer[:n])
		log.Printf("Received message: %s from %s", message, addr)

		if message == "ECHO-REQUEST" && *respond {
			response := []byte("ECHO-RESPONSE")
			_, err := conn.WriteToUDP(response, addr)
			if err != nil {
				log.Printf("Failed to send response: %v", err)
			} else {
				log.Printf("Sent response to %s", addr)
			}
		} else {
			log.Printf("Not responding to request")
		}

		// Add some simulated processing time
		time.Sleep(100 * time.Millisecond)
	}
}
