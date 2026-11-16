// Program to show the worker pool pattern: a fixed number of goroutines
// (workers) pick jobs from a shared channel and send results back.
package main

import (
	"fmt"
	"sync"
	"time"
)

type Result struct {
	WorkerID int
	Job      int
	Value    int
}

func worker(id int, jobs <-chan int, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		time.Sleep(time.Second) // simulate slow work
		results <- Result{WorkerID: id, Job: j, Value: j * 2}
	}
}

func main() {
	const numJobs = 5
	const numWorkers = 3

	jobs := make(chan int, numJobs)
	results := make(chan Result, numJobs)
	var wg sync.WaitGroup

	// Start the workers
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// Send the jobs and close the channel so workers stop when done
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// Close results once every worker has finished
	go func() {
		wg.Wait()
		close(results)
	}()

	for res := range results {
		fmt.Printf("Worker %d finished job %d, result: %d\n", res.WorkerID, res.Job, res.Value)
	}
}
