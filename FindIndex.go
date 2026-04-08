package main

import "fmt"

func twoSum(nums []int, target int) []int {
    lookup := make(map[int]int) // value -> index
    for i, num := range nums {
        complement := target - num
        if idx, found := lookup[complement]; found {
            return []int{idx, i}
        }
        lookup[num] = i
    }
    return nil
}

func main() {
    nums := []int{2, 7, 11, 15}
    target := 9
    result := twoSum(nums, target)
    fmt.Println("Indices:", result) // [0,1]
}
//Output
//Indices: [0 1]
