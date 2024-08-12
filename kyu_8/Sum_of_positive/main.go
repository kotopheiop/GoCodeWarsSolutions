package main

import "fmt"

func PositiveSum(numbers []int) int {
	var sum int

	for _, n := range numbers {
		if n > 0 {
			sum += n
		}
	}

	return sum
}

func main() {
	fmt.Println(PositiveSum([]int{1, 2, 3, 4, 5}))      // 15
	fmt.Println(PositiveSum([]int{-1, -2, -3, -4, -5})) // 0
	fmt.Println(PositiveSum([]int{1, -2, 3, 4, 5}))     // 13
}
