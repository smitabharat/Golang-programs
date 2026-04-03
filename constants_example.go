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
