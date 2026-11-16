// Program to print each vowel with its position in a string and then
// print the reversed string.
package main

import (
	"fmt"
	"strings"
)

func isVowel(ch rune) bool {
	return strings.ContainsRune("aeiouAEIOU", ch)
}

func main() {
	s := "hello"
	r := []rune(s)

	// Find vowels and their original positions (0 based index)
	for i, ch := range r {
		if isVowel(ch) {
			fmt.Printf("%c -> position %d\n", ch, i)
		}
	}

	// Reverse the complete string
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	fmt.Println("Reversed string:", string(r))
}

// Output:
// e -> position 1
// o -> position 4
// Reversed string: olleh
