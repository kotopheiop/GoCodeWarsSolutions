package main

import "fmt"

func FindMissingLetter(chars []rune) rune {
	for i := 0; i < len(chars)-1; i++ {
		if chars[i+1]-chars[i] > 1 {
			return chars[i] + 1
		}
	}
	return 0
}

func main() {
	fmt.Println(FindMissingLetter([]rune{'a', 'b', 'c', 'd', 'f'})) // 101
	fmt.Println(FindMissingLetter([]rune{'O', 'Q', 'R', 'S'}))      // 80
}
