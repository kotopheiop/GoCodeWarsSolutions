package main

import "fmt"

func monkeyCount(n int) []int {
	if n <= 0 {
		return nil
	}
	result := make([]int, n)
	for i := 1; i <= n; i++ {
		result[i-1] = i
	}

	return result
}

func main() {
	fmt.Println(monkeyCount(0)) // []
	fmt.Println(monkeyCount(6)) // [1 2 3 4 5 6]
}
