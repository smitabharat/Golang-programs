import (
	"fmt"
	"strings"
)

// write program to count of characters in a string
func main() {
	// Input string
	str := "aaabbbccdddee"

	// Create a map to store the frequency of characters
	charCount := make(map[string]int)

	// Split the string into individual characters
	splitStr := strings.Split(str, "")

	// Iterate through the split string and count each character
	for _, char := range splitStr {
		charCount[char]++
	}

	// Print the count of each character
	fmt.Println("Count of characters:")
	for char, count := range charCount {
		fmt.Printf("%s: %d\n", char, count)
	}

}
