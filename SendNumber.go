package main

import (
	"fmt"
	"sync"
)

func Send(numbers chan int, wg *sync.WaitGroup) {

	defer wg.Done()
	for i := 0; i <= 10; i++ {
		numbers <- i
	}
	close(numbers)

}
func main() {
	var wg sync.WaitGroup
	numbers := make(chan int)
	wg.Add(1)
	go Send(numbers, &wg)
	for num := range numbers {
		fmt.Println("numbers", num)
	}
	wg.Wait()
}
