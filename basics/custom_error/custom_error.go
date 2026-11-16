// Program to create and return a custom error type in Go.
package main

import (
	"errors"
	"fmt"
)

// CustomError is our own error type
type CustomError struct {
	Message string
}

// Error makes CustomError satisfy the built-in error interface
func (c *CustomError) Error() string {
	return c.Message
}

const minLength = 5

// ValidateInput returns an error when the input is shorter than minLength
func ValidateInput(input string) error {
	if len(input) < minLength {
		return &CustomError{
			Message: fmt.Sprintf("input string length should be at least %d characters", minLength),
		}
	}
	return nil
}

func main() {
	for _, input := range []string{"abc", "golang"} {
		err := ValidateInput(input)

		// errors.As checks whether err is (or wraps) a *CustomError
		var ce *CustomError
		if errors.As(err, &ce) {
			fmt.Printf("%q -> Error: %s\n", input, ce.Message)
			continue
		}
		fmt.Printf("%q -> Valid string\n", input)
	}
}

// Output:
// "abc" -> Error: input string length should be at least 5 characters
// "golang" -> Valid string
