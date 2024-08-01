package main

import "fmt"

func SmallestIntegerFinder(numbers []int) int {
	minValue := numbers[0]
	for _, n := range numbers {
		if n < minValue {
			minValue = n
		}
	}
	return minValue
}

func main() {
	fmt.Println(SmallestIntegerFinder([]int{1, 2, 3, 4, 5}))     // 1
	fmt.Println(SmallestIntegerFinder([]int{34, 15, 88, 2}))     // 2
	fmt.Println(SmallestIntegerFinder([]int{34, -345, -1, 100})) // -345
}
