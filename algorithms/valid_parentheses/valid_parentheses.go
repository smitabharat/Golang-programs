// Program to check whether a string of brackets is valid.
// Every opening bracket must be closed by the same type in the correct order.
package main

import "fmt"

func isValid(s string) bool {
	stack := []rune{}

	// map of closing -> opening bracket
	bracketMap := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, ch := range s {
		switch ch {
		case '(', '{', '[':
			// opening bracket -> push
			stack = append(stack, ch)
		case ')', '}', ']':
			// closing bracket -> top of stack must be the matching opener
			if len(stack) == 0 || stack[len(stack)-1] != bracketMap[ch] {
				return false
			}
			stack = stack[:len(stack)-1] // pop
		}
		// any other character is ignored
	}

	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("(){}{}")) // true
	fmt.Println(isValid("[{}}"))   // false
	fmt.Println(isValid("({}]"))   // false
	fmt.Println(isValid("{[()]}")) // true
}
