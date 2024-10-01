package main

import (
	"fmt"
	"strings"
)

func FindShort(s string) int {
	minElement := len(s)
	for _, str := range strings.Split(s, " ") {
		if len(str) < minElement {
			minElement = len(str)
		}
	}

	return minElement
}

func main() {
	fmt.Println(FindShort("bitcoin take over the world maybe who knows perhaps"))                // 3
	fmt.Println(FindShort("turns out random test cases are easier than writing out basic ones")) // 3
	fmt.Println(FindShort("Let's travel abroad shall we"))                                       // 2
}
