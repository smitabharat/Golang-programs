package main

import "fmt"

func isValid(s string) bool {
    stack := []rune{}

    // map of closing → opening
    bracketMap := map[rune]rune{
        ')': '(',
        '}': '{',
        ']': '[',
    }

    for _, ch := range s {

        // if closing bracket
        if open, ok := bracketMap[ch]; ok {
            if len(stack) == 0 || stack[len(stack)-1] != open {
                return false
            }
            stack = stack[:len(stack)-1] // pop
        } else {
            // opening bracket → push
            stack = append(stack, ch)
        }
    }

    return len(stack) == 0
}

func main() {
    fmt.Println(isValid("(){}{}")) // true
    fmt.Println(isValid("[{}}"))   // false
    fmt.Println(isValid("({}]"))   // false
}
