package main

import (
	"fmt"
	"strconv"
)

func Encode(str string, key int) []int {
	result := make([]int, len(str))
	keyStr := strconv.Itoa(key)

	for i, ch := range str {
		keyDigit := int(keyStr[i%len(keyStr)] - '0')
		result[i] = int(ch-'a'+1) + keyDigit
	}

	return result
}

func main() {
	fmt.Println(Encode("scout", 1939))       // 20, 12, 18, 30, 21
	fmt.Println(Encode("masterpiece", 1939)) // 14, 10, 22, 29, 6, 27, 19, 18, 6, 12, 8
}
