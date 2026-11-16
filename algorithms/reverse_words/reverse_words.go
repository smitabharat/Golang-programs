// Program to reverse each word of a sentence while keeping the word order.
// Input:  "this is input"
// Output: "siht si tupni"
package main

import (
	"fmt"
	"strings"
)

// reverseWord reverses a single word (rune based, so it is Unicode safe)
func reverseWord(word string) string {
	runes := []rune(word)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// reverseWords reverses each word in the input string
func reverseWords(input string) string {
	words := strings.Split(input, " ")
	for i, word := range words {
		words[i] = reverseWord(word)
	}
	return strings.Join(words, " ")
}

func main() {
	input := "this is input"
	output := reverseWords(input)
	fmt.Println("Reversed output:", output)
}

// Output:
// Reversed output: siht si tupni
