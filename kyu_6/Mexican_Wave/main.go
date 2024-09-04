package main

import (
	"fmt"
	"strings"
)

func wave(words string) []string {
	var result []string
	for i, char := range words {
		if char == ' ' {
			continue
		}
		waveWord := words[:i] + strings.ToUpper(string(char)) + words[i+1:]
		result = append(result, waveWord)
	}

	return result
}

func main() {
	fmt.Println(wave(" x yz"))     // []string{" X yz", " x Yz", " x yZ"}
	fmt.Println(wave("a a a a a")) // []string{"A a a a a", "a A a a a", "a a A a a", "a a a A a", "a a a a A"}
}
