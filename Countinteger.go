package main

import "fmt"

func main() {
    nums := []int{1, 2, 2, 3, 4, 3, 2, 5}

    // map to store counts
    counts := make(map[int]int)

    // counting occurrences
    for _, v := range nums {
        counts[v]++
    }

    // printing result
    for k, v := range counts {
        fmt.Printf("%d: %d\n", k, v)
    }
}
