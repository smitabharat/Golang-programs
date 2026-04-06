package main

import (
	"encoding/json"
	"net/http"
	"sync"
)

//write 2 gorutines for odd and even number and print them when i call an api

type Employee struct {
	Name    string
	Address string
}

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
	// employee := Employee{
	// 	Name:    "John",
	// 	Address: "Pune",
	// }
	numbers := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			numbers <- i
		}
		close(numbers)
	}()
	// Goroutine 2: Reads numbers from the channel and classifies them
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

	http.HandleFunc("/employee", func(w http.ResponseWriter, r *http.Request) {
		oddNumbers := []int{}
		evenNumbers := []int{}
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
