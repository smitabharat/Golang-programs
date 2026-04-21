package main

import "fmt"

func longestSubstring(s string) string {
    charMap := make(map[byte]int)
    left := 0
    maxLen := 0
    start := 0

    for right := 0; right < len(s); right++ {
        if idx, found := charMap[s[right]]; found && idx >= left {
            left = idx + 1
        }

        charMap[s[right]] = right

        if right-left+1 > maxLen {
            maxLen = right - left + 1
            start = left
        }
    }

    return s[start : start+maxLen]
}

func main() {
    fmt.Println(longestSubstring("abccddgh"))
}
