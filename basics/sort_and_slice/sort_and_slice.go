// Program to generate random integers, sort them and print the elements
// from index 2 (the third element) to the end.
package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	// Since Go 1.20 the global random generator is seeded automatically,
	// so rand.Seed (now deprecated) is not needed.
	arr := make([]int, 5)
	for i := range arr {
		arr[i] = rand.Intn(100) // random number between 0 and 99
	}
	fmt.Println("Unsorted list of integers:", arr)

	sort.Ints(arr)
	fmt.Println("Sorted list of integers:", arr)

	// Indexes start at 0, so index 2 is the third element
	if len(arr) > 2 {
		fmt.Println("Elements from index 2 to last:", arr[2:])
	} else {
		fmt.Println("Not enough elements to print from index 2")
	}
}
