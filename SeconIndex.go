package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	// Generate a random integer list
	rand.Seed(time.Now().UnixNano()) // Fixing the call
	arr := make([]int, 5)
	for i := 0; i < 5; i++ {
		arr[i] = rand.Intn(100) // Generate random numbers between 0 and 99
	}

	fmt.Println("Unsorted list of integers:", arr)

	// Sort in ascending order
	sort.Ints(arr)

	fmt.Println("Sorted list of integers:", arr)

	// Print elements from second index to last element
	if len(arr) > 2 {
		fmt.Println("Elements from second index to last:", arr[2:])
	} else {
		fmt.Println("Not enough elements to print from second index to last")
	}
}
