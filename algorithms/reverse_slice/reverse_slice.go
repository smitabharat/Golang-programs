// Program to return a reversed copy of a slice without changing the original.
// Input: [1 2 3 4]  Output: [4 3 2 1]
package main

import "fmt"

func reversedCopy(arr []int) []int {
	n := len(arr)
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = arr[n-1-i]
	}
	return result
}

func main() {
	arr := []int{1, 2, 3, 4}
	rev := reversedCopy(arr)
	fmt.Println("Original:", arr)
	fmt.Println("Reversed:", rev)
}

// Output:
// Original: [1 2 3 4]
// Reversed: [4 3 2 1]
