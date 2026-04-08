// Function to reverse each word in the input string
func reverseWords(input string) string {
	words := strings.Split(input, " ")
	for i, word := range words {
		words[i] = reverseWord(word)
	}
	return strings.Join(words, " ")
}

func main() {
	input := "this is input"
	output := reverseWords(input)
	fmt.Println("Reversed output:", output)
}
