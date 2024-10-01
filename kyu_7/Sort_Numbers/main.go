package main

import (
	"fmt"
	"sort"
)

func SortNumbers(numbers []int) []int {
	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] < numbers[j]
	})

	return numbers
}

func main() {
	fmt.Println(SortNumbers([]int{1, 2, 10, 50, 5})) // [1,2,5,10,50]
	fmt.Println(SortNumbers([]int{}))                // []
}
