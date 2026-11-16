// Program to count how many times each character appears in a string.
package main

import (
	"fmt"
	"sort"
)

func main() {
	str := "aaabbbccdddee"

	// Map to store the frequency of each character
	charCount := make(map[rune]int)
	for _, ch := range str {
		charCount[ch]++
	}

	// Map iteration order is random in Go, so sort the keys for stable output
	keys := make([]rune, 0, len(charCount))
	for ch := range charCount {
		keys = append(keys, ch)
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

	fmt.Println("Count of characters:")
	for _, ch := range keys {
		fmt.Printf("%c: %d\n", ch, charCount[ch])
	}
}

// Output:
// Count of characters:
// a: 3
// b: 3
// c: 2
// d: 3
// e: 2
