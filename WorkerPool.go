package main

import (
	"fmt"
	"sync"
	"time"
)

func Worker(id int, job <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range job {
		time.Sleep(time.Second)
		results <- j * 2
	}
}

func main() {
	const numJobs = 5
	const numWorkers = 3

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	var wg sync.WaitGroup

	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go Worker(w, jobs, results, &wg)

	}
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()
	for res := range results {
		fmt.Printf("Result Received:%d\n", res)
	}
}
