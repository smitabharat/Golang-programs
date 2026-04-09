package main

import (
	"fmt"
)

func main() {
	a := []int{3, 1, 4, 3, 2, 7, 8, 9, 0}
	a = a[:0]
	fmt.Println(len(a), cap(a), a)     //o/p -  0 9 []

	x := []int{3, 1, 4, 3, 2, 7, 8, 9, 0}
	b := x[2:4]
	fmt.Println(len(b), cap(b), b)        //o/p - 2 7 [4 3]
	s := append(b, 34, 21)
	fmt.Println(len(s), cap(s), s)   //o/p -  4 7 [4 3 34 21]
	s = s[1:6]
	fmt.Println(len(s), cap(s), s)   // o/p - 5 6 [3 34 21 8 9]

	y := []int{1, 4, 2, 6, 7, 9, 2, 1, 0, 1}
	y = y[2:5]
	fmt.Println(len(y), cap(y), y)   // o/p -  3 8 [2 6 7]
	y = y[2:8]
	fmt.Println(len(y), cap(y), y)   //o/p -  6 6 [7 9 2 1 0 1]
	y = y[2:]
	fmt.Println(len(y), cap(y), y)    // o/p - 4 4 [2 1 0 1]
}
