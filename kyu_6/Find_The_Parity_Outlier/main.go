package main

import "fmt"

func FindOdd(seq []int) int {
	m := make(map[int]int)
	odd := 0
	for _, val := range seq {
		m[val] += 1
	}

	for key, val := range m {
		if val%2 != 0 {
			odd = key
		}
	}
	return odd
}

func main() {
	fmt.Println(FindOdd([]int{1, 2, 3, 4, 5}))
	fmt.Println(FindOdd([]int{1, 2, 2, 3, 3, 3, 4, 3, 3, 3, 2, 2, 1}))
}
