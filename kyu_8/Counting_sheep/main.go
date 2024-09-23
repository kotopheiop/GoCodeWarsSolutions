package main

import "fmt"

func CountSheeps(numbers []bool) int {
	var result int

	for _, val := range numbers {
		if val {
			result++
		}
	}

	return result
}

func main() {
	fmt.Println(CountSheeps([]bool{true, false})) // 1
	fmt.Println(CountSheeps([]bool{               //17
		true, true, true, false,
		true, true, true, true,
		true, false, true, false,
		true, false, false, true,
		true, true, true, true,
		false, false, true, true,
	}))
}
