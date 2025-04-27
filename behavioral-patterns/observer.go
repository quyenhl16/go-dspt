package main

import "fmt"

// Observer is an interface that defines the Update method.
// All observers must implement this to receive notifications.
type Observer interface {
	Update(string)
}

// Subject defines the methods to manage observers.
// It allows registering, deregistering, and notifying observers.
type Subject interface {
	Register(Observer)
	Deregister(Observer)
	NotifyAll()
}

// NewsAgency is the concrete subject.
// It holds a list of observers and the current news.
type NewsAgency struct {
	observers []Observer
	news      string
}

// Register adds an observer to the list.
func (n *NewsAgency) Register(o Observer) {
	n.observers = append(n.observers, o)
}

// Deregister removes an observer from the list.
func (n *NewsAgency) Deregister(o Observer) {
	for i, observer := range n.observers {
		if observer == o {
			// Remove the observer from the slice.
			n.observers = append(n.observers[:i], n.observers[i+1:]...)
			break
		}
	}
}

// NotifyAll informs all registered observers about the update.
func (n *NewsAgency) NotifyAll() {
	for _, observer := range n.observers {
		observer.Update(n.news) // Call each observer's Update method.
	}
}

// AddNews sets new news and notifies all observers.
func (n *NewsAgency) AddNews(news string) {
	n.news = news
	n.NotifyAll()
}

// Subscriber is a concrete implementation of the Observer interface.
// It represents a user who wants to receive updates.
type Subscriber struct {
	name string
}

// Update is called by the subject to notify this subscriber.
func (s *Subscriber) Update(news string) {
	fmt.Printf("%s received news: %s\n", s.name, news)
}

func main() {
	// Create a NewsAgency (Subject)
	agency := &NewsAgency{}

	// Create Subscribers (Observers)
	sub1 := &Subscriber{name: "Alice"}
	sub2 := &Subscriber{name: "Bob"}

	// Register subscribers to the agency
	agency.Register(sub1)
	agency.Register(sub2)

	// Add news - both Alice and Bob will receive it
	agency.AddNews("New Go release is out!")

	// Deregister Alice
	agency.Deregister(sub1)

	// Add another news - only Bob will receive it
	agency.AddNews("Golang 2.0 announced!")
}
