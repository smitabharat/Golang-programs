// Two goroutines: one generates numbers 1 to 10, the other classifies them
// as odd or even. The result is served at GET http://localhost:8080/numbers
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

type NumberClassification struct {
	OddNumbers  []int `json:"odd_numbers"`
	EvenNumbers []int `json:"even_numbers"`
}

func classify(max int) NumberClassification {
	numbers := make(chan int)
	var result NumberClassification
	var wg sync.WaitGroup

	// Goroutine 1: generate numbers
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= max; i++ {
			numbers <- i
		}
		close(numbers)
	}()

	// Goroutine 2: classify numbers (only this goroutine writes to result)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for num := range numbers {
			if num%2 == 0 {
				result.EvenNumbers = append(result.EvenNumbers, num)
			} else {
				result.OddNumbers = append(result.OddNumbers, num)
			}
		}
	}()

	wg.Wait()
	return result
}

func main() {
	http.HandleFunc("/numbers", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(classify(10)); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// GET /numbers response:
// {"odd_numbers":[1,3,5,7,9],"even_numbers":[2,4,6,8,10]}
