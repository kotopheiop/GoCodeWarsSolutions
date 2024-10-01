package main

import (
	"fmt"
	"strings"
)

func High(s string) string {
	words := strings.Split(s, " ")
	highestWord := ""
	maxScore := 0

	for _, word := range words {
		currentScore := 0

		for _, ch := range word {
			// Позиция буквы в алфавите: 'a' = 1, 'b' = 2, ..., 'z' = 26
			currentScore += int(ch - 'a' + 1)
		}

		if currentScore > maxScore {
			maxScore = currentScore
			highestWord = word
		}
	}

	return highestWord
}

func main() {
	fmt.Println(High("man i need a taxi up to ubud"))             // taxi
	fmt.Println(High("what time are we climbing up the volcano")) // volcano
	fmt.Println(High("take me to semynak"))                       // semynak
}
