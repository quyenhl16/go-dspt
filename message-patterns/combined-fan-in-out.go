package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Job defines the structure of a job
type Job struct {
	ID    int
	Value int
}

// Result defines the result after processing a job
type Result struct {
	WorkerID int
	JobID    int
	Output   int
}

// cGenerator produces jobs and sends them to a shared channel
func cGenerator(ctx context.Context, total int) <-chan Job {
	out := make(chan Job)

	go func() {
		defer close(out)
		for i := 1; i <= total; i++ {
			select {
			case <-ctx.Done():
				fmt.Println("Generator stopped")
				return
			case out <- Job{ID: i, Value: rand.Intn(100)}:
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	return out
}

// cWorker processes jobs and sends results to the output channel
func cWorker(ctx context.Context, id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d stopped\n", id)
			return
		case job, ok := <-jobs:
			if !ok {
				return
			}
			time.Sleep(300 * time.Millisecond) // simulate work
			result := Result{
				WorkerID: id,
				JobID:    job.ID,
				Output:   job.Value * 2, // example processing: multiply by 2
			}
			select {
			case <-ctx.Done():
				return
			case results <- result:
			}
		}
	}
}

// cFanIn collects all results into one channel
func cFanIn(ctx context.Context, numWorkers int, results <-chan Result) <-chan Result {
	out := make(chan Result)

	go func() {
		defer close(out)
		for {
			select {
			case <-ctx.Done():
				return
			case res, ok := <-results:
				if !ok {
					return
				}
				out <- res
			}
		}
	}()

	return out
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Create job source (fan-out input)
	jobs := cGenerator(ctx, 20)

	// Shared result channel for fan-in
	results := make(chan Result)

	// Fan-out: Start workers
	numWorkers := 3
	var wg sync.WaitGroup
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go cWorker(ctx, i, jobs, results, &wg)
	}

	// Fan-in: Collect results from all workers
	go func() {
		wg.Wait()
		close(results)
	}()

	// Read from the fan-in output
	for res := range cFanIn(ctx, numWorkers, results) {
		fmt.Printf("Result: Job %d processed by Worker %d -> Output: %d\n", res.JobID, res.WorkerID, res.Output)
	}

	fmt.Println("Pipeline complete")
}
