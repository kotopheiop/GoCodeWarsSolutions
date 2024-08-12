package main

import "fmt"

func Solution(str string) []string {
	var result []string

	for i := 0; i < len(str); i += 2 {
		end := i + 2
		if end > len(str) {
			end = len(str)
		}
		pair := str[i:end]

		if len(pair) < 2 {
			pair += "_"
		}

		result = append(result, pair)
	}

	return result
}

func main() {
	fmt.Println(Solution("abcd"))  // [ab cd]
	fmt.Println(Solution("abcde")) // [ab cd e_]
}
