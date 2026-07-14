package main

import "fmt"

func main() {
	input := "gogolang"

	// Count frequency of each character
	freq := make(map[rune]int)

	for _, ch := range input {
		freq[ch]++
	}

	// Find first non-repeating character
	for _, ch := range input {
		if freq[ch] == 1 {
			fmt.Printf("%c\n", ch)
			return
		}
	}

	fmt.Println("No non-repeating character found")
}
