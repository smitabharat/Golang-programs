package main

import (
	"fmt"
)

func main() {
	a := []int{3, 1, 4, 3, 2, 7, 8, 9, 0}
	a = a[:0]
	fmt.Println(len(a), cap(a), a)     //o/p -  0 9 []   
	//a = a[:0]
   //Start = 0, End = 0
   //Length = 0
  //Capacity = still from original backing array

	x := []int{3, 1, 4, 3, 2, 7, 8, 9, 0}
	b := x[2:4]
	fmt.Println(len(b), cap(b), b)        //o/p - 2 7 [4 3]
	//b = x[2:4] → [4, 3]
    //Length = 4 - 2 = 2
    //Capacity = from index 2 → end of array
    //cap = len(x) - 2 = 9 - 2 = 7
	s := append(b, 34, 21)
	fmt.Println(len(s), cap(s), s)   //o/p -  4 7 [4 3 34 21]
	//b has cap = 7
   //Current len = 2
   //Adding 2 elements → total = 4
	s = s[1:6]
	fmt.Println(len(s), cap(s), s)   // o/p - 5 6 [3 34 21 8 9]
   //New slice: [3, 34, 21, 7, 8]
  //Length = 6 - 1 = 5
  //Capacity = old cap - 1 = 7 - 1 = 6
	y := []int{1, 4, 2, 6, 7, 9, 2, 1, 0, 1}
	y = y[2:5]
	fmt.Println(len(y), cap(y), y)   // o/p -  3 8 [2 6 7]
	//[2, 6, 7]
   //len = 3
  //cap = 10 - 2 = 8
	y = y[2:8]
	fmt.Println(len(y), cap(y), y)   //o/p -  6 6 [7 9 2 1 0 1]
	//[7, 9, 2, 1, 0, 1]
   //len = 8 - 2 = 6
   //cap = 6
	y = y[2:]
	fmt.Println(len(y), cap(y), y)    // o/p - 4 4 [2 1 0 1]
	//[2, 1, 0, 1]
    //len = 4
   //cap = 4
}
