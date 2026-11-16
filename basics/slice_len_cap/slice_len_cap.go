// Program to understand how len() and cap() change when you re-slice
// and append to a Go slice.
package main

import "fmt"

func main() {
	// 1. Re-slice to zero length
	a := []int{3, 1, 4, 3, 2, 7, 8, 9, 0}
	a = a[:0]
	fmt.Println(len(a), cap(a), a) // 0 9 []
	// start = 0, end = 0 -> len = 0
	// cap stays 9 because the backing array is unchanged

	// 2. Slice in the middle
	x := []int{3, 1, 4, 3, 2, 7, 8, 9, 0}
	b := x[2:4]
	fmt.Println(len(b), cap(b), b) // 2 7 [4 3]
	// len = 4 - 2 = 2
	// cap = len(x) - 2 = 9 - 2 = 7

	// 3. Append within capacity reuses (and overwrites) the backing array
	s := append(b, 34, 21)
	fmt.Println(len(s), cap(s), s) // 4 7 [4 3 34 21]
	// cap is 7, so no new array is allocated.
	// x is now [3 1 4 3 34 21 8 9 0] (x[4] and x[5] were overwritten)

	// 4. Re-slice beyond len (allowed up to cap)
	s = s[1:6]
	fmt.Println(len(s), cap(s), s) // 5 6 [3 34 21 8 9]
	// len = 6 - 1 = 5
	// cap = 7 - 1 = 6
	// elements 8 and 9 come from the original backing array x

	// 5. Re-slicing a re-slice
	y := []int{1, 4, 2, 6, 7, 9, 2, 1, 0, 1}
	y = y[2:5]
	fmt.Println(len(y), cap(y), y) // 3 8 [2 6 7]
	// len = 5 - 2 = 3, cap = 10 - 2 = 8

	y = y[2:8]
	fmt.Println(len(y), cap(y), y) // 6 6 [7 9 2 1 0 1]
	// len = 8 - 2 = 6, cap = 8 - 2 = 6

	y = y[2:]
	fmt.Println(len(y), cap(y), y) // 4 4 [2 1 0 1]
	// len = 6 - 2 = 4, cap = 6 - 2 = 4
}
