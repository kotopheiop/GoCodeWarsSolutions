package main

import (
	"fmt"
	"strings"
)

func SpinWords(str string) string {
	var res string
	if len(str) == 0 {
		return res
	}
	splitStr := strings.Split(str, " ")
	for _, word := range splitStr {
		if len(word) >= 5 {
			word = Reverse(word)
		}
		res = res + word + " "
	}

	return strings.TrimSpace(res)
}

func Reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}

	return
}

func main() {
	fmt.Println(SpinWords("Hey fellow warriors")) // Hey wollef sroirraw
	fmt.Println(SpinWords("Welcome"))             // emocleW
	fmt.Println(SpinWords("CodeWars"))            // sraWedoC

}
