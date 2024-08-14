package main

import (
	"fmt"
	"strconv"
)

func Digitize(n int) []int {
	var result []int

	nReversString := reverse(strconv.Itoa(n))
	for _, digit := range nReversString {
		digitInt, _ := strconv.Atoi(string(digit))
		result = append(result, digitInt)
	}

	return result
}

func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}

	return
}

func main() {
	fmt.Println(Digitize(35231)) // [1 3 2 5 3]
	fmt.Println(Digitize(0))     // [0]
}
