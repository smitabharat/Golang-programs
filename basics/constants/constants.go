// In a const block, a constant with no value repeats the previous
// expression, so b and c both get the value 1.
package main

import "fmt"

const (
	a = 1
	b
	c
)

func main() {
	fmt.Println(a, b, c) // Output: 1 1 1
}
