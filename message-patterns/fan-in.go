package main

import (
	"context"
	"fmt"
	"time"
)

// producer generates messages and sends them into its channel until the context is done.
func producer(ctx context.Context, label string, interval time.Duration) <-chan string {
	ch := make(chan string)

	go func() {
		defer close(ch) // Ensure the channel is closed when the goroutine exits

		for i := 0; ; i++ {
			select {
			case <-ctx.Done():
				// Exit the goroutine when the context is cancelled
				fmt.Println(label, "stopped")
				return
			case ch <- fmt.Sprintf("%s: %d", label, i):
				time.Sleep(interval) // Simulate work by sleeping
			}
		}
	}()

	return ch
}

// cFanIn merges multiple input channels into a single output channel,
// and stops when the context is cancelled.
func fanIn(ctx context.Context, channels ...<-chan string) <-chan string {
	out := make(chan string)

	// Start a goroutine for each input channel
	for _, ch := range channels {
		go func(c <-chan string) {
			for {
				select {
				case <-ctx.Done():
					// Exit goroutine when context is cancelled
					return
				case msg, ok := <-c:
					if !ok {
						// If the input channel is closed, exit
						return
					}
					out <- msg
				}
			}
		}(ch)
	}

	return out
}

func main() {
	// Create a context with timeout to stop everything after 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel() // Ensure context is cancelled on exit

	// Start two producers
	ch1 := producer(ctx, "Producer A", 500*time.Millisecond)
	ch2 := producer(ctx, "Producer B", 700*time.Millisecond)

	// Merge the channels using fan-in pattern
	merged := fanIn(ctx, ch1, ch2)

	// Read from merged channel until context is done
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Main: context cancelled, shutting down")
			return
		case msg := <-merged:
			fmt.Println("Received:", msg)
		}
	}
}
