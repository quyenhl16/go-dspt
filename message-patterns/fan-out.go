package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// cGenerator sends integers to a channel until context is canceled
func generator(ctx context.Context) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Generator stopped")
				return
			case out <- rand.Intn(100): // Simulate generating work
				time.Sleep(300 * time.Millisecond)
			}
		}
	}()

	return out
}

// cWorker reads from input channel, processes the job, and prints the result
func worker(ctx context.Context, id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d stopped\n", id)
			return
		case job, ok := <-jobs:
			if !ok {
				// Channel closed, exit
				fmt.Printf("Worker %d: job channel closed\n", id)
				return
			}
			// Simulate work
			time.Sleep(500 * time.Millisecond)
			fmt.Printf("Worker %d processed job: %d\n", id, job)
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Create job source
	jobs := generator(ctx)

	// Fan-out: Start multiple workers consuming from the same job channel
	var wg sync.WaitGroup
	numWorkers := 3
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(ctx, i, jobs, &wg)
	}

	// Wait for all workers to finish
	wg.Wait()
	fmt.Println("All workers done")
}
