package main

import (
	"fmt"
	"unicode"
)

func ReverseWords(str string) string {
	runes := []rune(str)
	start := 0
	for i := 0; i < len(runes); i++ {
		if unicode.IsSpace(runes[i]) {
			reverse(runes[start:i])
			start = i + 1
		}
	}
	reverse(runes[start:])
	return string(runes)
}

func reverse(word []rune) {
	for i, j := 0, len(word)-1; i < j; i, j = i+1, j-1 {
		word[i], word[j] = word[j], word[i]
	}
}

func main() {
	fmt.Println(ReverseWords("This is an example!"))
	fmt.Printf(ReverseWords("double  spaces"))
}
