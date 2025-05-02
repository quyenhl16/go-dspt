package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	maxQueueSize     = 3
	numProducers     = 2
	numConsumers     = 2
	itemsPerProducer = 5
)

var (
	queue []int          // Shared bounded queue
	mu    = sync.Mutex{} // Mutex for safe access
	cond  = sync.NewCond(&mu)
	wg    sync.WaitGroup // WaitGroup to wait for all producers/consumers
)

// Producer goroutine
func producer(id int) {
	defer wg.Done() // Notify WaitGroup when done

	for i := 0; i < itemsPerProducer; i++ {
		time.Sleep(time.Millisecond * 300) // simulate work

		cond.L.Lock()
		for len(queue) == maxQueueSize {
			fmt.Printf("[Producer %d] Queue full, waiting...\n", id)
			cond.Wait() // Wait until not full
		}

		item := id*100 + i // Unique item
		queue = append(queue, item)
		fmt.Printf("[Producer %d] Produced: %d\n", id, item)

		cond.Broadcast() // Notify all consumers that something changed
		cond.L.Unlock()
	}
}

// Consumer goroutine
func consumer(id int) {
	defer wg.Done() // Notify WaitGroup when done

	for {
		cond.L.Lock()

		for len(queue) == 0 {
			if wgWaitDone() {
				cond.L.Unlock()
				return // Exit if no more producers will add
			}
			fmt.Printf("[Consumer %d] Queue empty, waiting...\n", id)
			cond.Wait() // Wait until not empty
		}

		item := queue[0]
		queue = queue[1:]
		fmt.Printf("[Consumer %d] Consumed: %d\n", id, item)

		cond.Broadcast() // Notify all producers that space is available
		cond.L.Unlock()

		time.Sleep(time.Millisecond * 500) // simulate work
	}
}

// wgWaitDone checks if all producers are done AND queue is empty
func wgWaitDone() bool {
	// Peek inside WaitGroup count (not directly accessible, so a workaround is needed in real apps)
	// For simplicity here, we check if queue is empty and sleep count time
	time.Sleep(100 * time.Millisecond)
	return len(queue) == 0
}

func main() {
	// Start multiple producers
	for i := 1; i <= numProducers; i++ {
		wg.Add(1)
		go producer(i)
	}

	// Start multiple consumers
	for i := 1; i <= numConsumers; i++ {
		wg.Add(1)
		go consumer(i)
	}

	// Wait for all producers and consumers to finish
	wg.Wait()
	fmt.Println("\n✅ All producers and consumers have completed.")
}
