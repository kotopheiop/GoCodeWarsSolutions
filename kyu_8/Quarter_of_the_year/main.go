package main

import "fmt"

func QuarterOf(month int) int {
	return (month + 2) / 3
}

func main() {
	fmt.Println(QuarterOf(1))  // 1
	fmt.Println(QuarterOf(12)) // 4
}
