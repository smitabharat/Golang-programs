// Program to show how panic stops normal execution and how recover
// can catch a panic so the program keeps running.
package main

import "fmt"

func divide(a, b int) (result int, err error) {
	// recover() only works inside a deferred function
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("recovered from panic: %v", r)
		}
	}()

	if b == 0 {
		panic("division by zero")
		// Any code written here would never run (unreachable)
	}
	return a / b, nil
}

func main() {
	if res, err := divide(10, 2); err == nil {
		fmt.Println("10 / 2 =", res)
	}

	if _, err := divide(10, 0); err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("Program continues after recover")
}

// Output:
// 10 / 2 = 5
// Error: recovered from panic: division by zero
// Program continues after recover
