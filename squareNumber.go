package main

import "fmt"

func squareNumber(n int)int {
	if n < 0 {
	  n = -n
	}

	if n == 0 {
	  return 0
	}

	return squareNumber(n-1) + (2*n - 1)
}

func main() {
	number := 3
	result := squareNumber(number)
	fmt.Printf("The square of %d is %d\n", number, result)
}
