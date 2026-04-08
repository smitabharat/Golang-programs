package main

import (
	"fmt"
	"sync"
)

func printEven(max int, evenCh, oddCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range evenCh {
		if num > max {
			close(oddCh)
			return
		}
		fmt.Printf("Even: %d\n", num)
		oddCh <- num + 1
	}
}

func printOdd(max int, oddCh, evenCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range oddCh {
		if num > max {
			close(evenCh)
			return
		}
		fmt.Printf("Odd: %d\n", num)
		evenCh <- num + 1
	}
}

func main() {
	var wg sync.WaitGroup
	max := 10

	oddCh := make(chan int)
	evenCh := make(chan int)

	wg.Add(2)
	go printEven(max, evenCh, oddCh, &wg)
	go printOdd(max, oddCh, evenCh, &wg)

	// Start the sequence with the first odd number
	oddCh <- 1

	wg.Wait()
}
