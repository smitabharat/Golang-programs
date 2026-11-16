// Program to split a comma separated string, sort the words alphabetically
// and print each one with its first letter in upper case.
package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fruits := "banana,apple,orange,grape"

	fruitList := strings.Split(fruits, ",")
	for i, fruit := range fruitList {
		fruitList[i] = strings.TrimSpace(fruit)
	}

	sort.Strings(fruitList)

	for _, fruit := range fruitList {
		if fruit == "" {
			continue // avoid panic on fruit[:1] for empty values like "a,,b"
		}
		fmt.Println(strings.ToUpper(fruit[:1]) + strings.ToLower(fruit[1:]))
	}
}

// Output:
// Apple
// Banana
// Grape
// Orange
