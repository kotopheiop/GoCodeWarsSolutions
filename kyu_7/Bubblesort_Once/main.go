package main

import "fmt"

func BubblesortOnce(numbers []int) []int {
	// Create a copy of the original slice
	numbersCopy := make([]int, len(numbers))
	copy(numbersCopy, numbers)

	for i := 0; i < len(numbersCopy)-1; i++ {
		if numbersCopy[i] > numbersCopy[i+1] {
			numbersCopy[i], numbersCopy[i+1] = numbersCopy[i+1], numbersCopy[i]
		}
	}
	return numbersCopy
}

func main() {
	fmt.Println(BubblesortOnce([]int{9, 7, 5, 3, 1, 2, 4, 6, 8}))
}
