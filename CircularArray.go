package main

import "fmt"

func nextGreaterNumber(arr []int) []int {
	n := len(arr)
	result := make([]int, n)

	for i := range result {
		result[i] = -1
	}

	stack := []int{}

	for i := 0; i < 2*n; i++ {
		num := arr[i%n]

		for len(stack) > 0 && num > arr[stack[len(stack)-1]] {
			idx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[idx] = num
		}

		if i < n {
			stack = append(stack, i)
		}
	}

	return result
}

func main() {
	arr := []int{1, 2, 1}
	fmt.Println(nextGreaterNumber(arr))
}
