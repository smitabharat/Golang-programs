package main

import (
	"fmt"
)

// Custom Error Struct
type CustomError struct {
	Message string
}

// Implement Error() method
func (c *CustomError) Error() string {
	return c.Message
}

// Function taking string input
func InputStr(input string) error {

	if len(input) < 5 {
		return &CustomError{
			Message: "input string length should be greater than 5",
		}
	}

	return nil
}

func main() {

	err := InputStr("abc")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Valid String")
}

