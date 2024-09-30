package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Points(results []string) int {
	points := 0

	for _, result := range results {
		scores := strings.Split(result, ":")
		x, _ := strconv.Atoi(scores[0])
		y, _ := strconv.Atoi(scores[1])

		if x > y {
			points += 3 // Win
		} else if x < y {
			points += 0 // Loss
		} else {
			points += 1 // Tie
		}
	}

	return points
}

func main() {
	fmt.Println(Points([]string{"1:0", "2:0", "3:0", "4:0", "2:1", "3:1", "4:1", "3:2", "4:2", "4:3"})) // 30
	fmt.Println(Points([]string{"0:1", "0:2", "0:3", "0:4", "1:2", "1:3", "1:4", "2:3", "2:4", "3:4"})) // 0
	fmt.Println(Points([]string{"1:0", "2:0", "3:0", "4:0", "2:1", "1:3", "1:4", "2:3", "2:4", "3:4"})) // 15
}
