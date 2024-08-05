package main

import (
	"fmt"
	"strings"
)

func AbbrevName(name string) string {
	var result string
	words := strings.Split(name, " ")

	for _, word := range words {
		result += string(word[0]) + "."
	}

	return strings.ToUpper(result[:3])
}

func main() {
	fmt.Println(AbbrevName("Sam Harris"))     // S.H
	fmt.Println(AbbrevName("patrick feeney")) // P.F
}
