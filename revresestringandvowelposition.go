package main

import "fmt"

func main() {
	s := "hello"

	r := []rune(s)

	// Find vowels and their original positions
	for i, ch := range r {
		if isVowel(ch) {
			fmt.Printf("%c → position %d\n", ch, i)
		}
	}

	// Reverse complete string
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	fmt.Println("Reversed string:", string(r))
}

func isVowel(ch rune) bool {
	return ch == 'a' || ch == 'e' || ch == 'i' ||
		ch == 'o' || ch == 'u'
}
