package main

import (
	"fmt"
	"strconv"
)

func SumMix(arr []any) int {
	var result int
	for _, v := range arr {
		switch v.(type) {
		case int:
			result += v.(int)
			continue
		case string:
			vInt, _ := strconv.Atoi(v.(string))
			result += vInt
			continue
		}
	}

	return result
}

func main() {
	fmt.Println(SumMix([]any{9, 3, "7", "3"}))                  // 22
	fmt.Println(SumMix([]any{"5", "0", 9, 3, 2, 1, "9", 6, 7})) // 42
}
