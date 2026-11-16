// Program to count how many times each integer appears in a slice.
package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, 2, 3, 4, 3, 2, 5}

	// Map to store counts
	counts := make(map[int]int)
	for _, v := range nums {
		counts[v]++
	}

	// Map iteration order is random, so sort the keys for stable output
	keys := make([]int, 0, len(counts))
	for k := range counts {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		fmt.Printf("%d: %d\n", k, counts[k])
	}
}

// Output:
// 1: 1
// 2: 3
// 3: 2
// 4: 1
// 5: 1
