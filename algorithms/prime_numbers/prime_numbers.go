// Program to print all prime numbers from 1 to 100.
package main

import "fmt"

// isPrime checks divisors only up to the square root of num
func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("Prime numbers from 1 to 100 are:")
	for i := 1; i <= 100; i++ {
		if isPrime(i) {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}
