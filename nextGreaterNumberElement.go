package main

import "fmt"

func nextGreaterNumber(arr []int) []int {
	n := len(arr)
	result := make([]int, n)

	for i := range result {
		result[i] = -1
	}
	stack := []int{}
	for i := 0; i < n; i++ {
		for len(stack) > 0 && arr[i] > arr[stack[len(stack)-1]] {
			idx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[idx] = arr[i]
		}
		stack = append(stack, i)
	}
	return result
}

func main() {
	arr := []int{4, 5, 2, 25}
	fmt.Println(nextGreaterNumber(arr))
}
