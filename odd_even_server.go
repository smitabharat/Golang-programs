package main

import (
	"encoding/json"
	"net/http"
	"sync"
)

type NumberClassification struct {
	OddNumbers  []int `json:"odd_numbers"`
	EvenNumbers []int `json:"even_numbers"`
}

var (
	oddNumbers  []int
	evenNumbers []int
	mu          sync.Mutex
)

func main() {
	numbers := make(chan int)
	var wg sync.WaitGroup

	// Goroutine 1: Generate numbers
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			numbers <- i
		}
		close(numbers)
	}()

	// Goroutine 2: Classify numbers
	wg.Add(1)
	go func() {
		defer wg.Done()
		for num := range numbers {
			mu.Lock()
			if num%2 == 0 {
				evenNumbers = append(evenNumbers, num)
			} else {
				oddNumbers = append(oddNumbers, num)
			}
			mu.Unlock()
		}
	}()

	wg.Wait()

	http.HandleFunc("/numbers", func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		response := NumberClassification{
			OddNumbers:  oddNumbers,
			EvenNumbers: evenNumbers,
		}
		mu.Unlock()

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	http.ListenAndServe(":8080", nil)
}
