package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	ans := twoSum(nums, target)
	fmt.Println(ans)
}

func twoSum(nums []int, target int) []int {
	indexMap := make(map[int]int)

	for i, num := range nums {
		complement := target - num

		if idx, exists := indexMap[complement]; exists {
			return []int{idx, i}
		}

		indexMap[num] = i
	}
	return nil
}
