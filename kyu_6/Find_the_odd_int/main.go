package main

import "fmt"

func FindOutlier(integers []int) int {
	odds := make([]int, 0)
	evens := make([]int, 0)

	for _, value := range integers {
		if value%2 == 0 {
			evens = append(evens, value)
		} else {
			odds = append(odds, value)
		}
	}

	if len(odds) < len(evens) {
		return odds[0]
	} else {
		return evens[0]
	}
}

func main() {
	fmt.Println(FindOutlier([]int{2, 4, 0, 100, 4, 11, 2602, 36}))
	fmt.Println(FindOutlier([]int{160, 3, 1719, 19, 11, 13, -21}))
}
