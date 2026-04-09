package main

import (
	"fmt"
	"sync"
)

func sumOfSquaresOfEvens(numbers []int) int {
	var wg sync.WaitGroup
	sumChan := make(chan int)

	for _, num := range numbers {
		if num%2 == 0 { // only even numbers
			wg.Add(1)
			go func(n int) {
				defer wg.Done()
				sumChan <- n * n // square and send to channel
			}(num)
		}
	}

	// closing channel when all goroutines finish
	go func() {
		wg.Wait()
		close(sumChan)
	}()

	total := 0
	for val := range sumChan {
		total += val
	}

	return total
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := sumOfSquaresOfEvens(nums)
	fmt.Println("Sum of squares of even numbers:", result)
}
