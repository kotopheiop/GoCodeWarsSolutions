package main

import "fmt"

func Summation(n int) int {
	var result int
	for i := 1; i <= n; i++ {
		result += i
	}

	return result
}

func main() {
	fmt.Println(Summation(2))
	fmt.Println(Summation(8))
	fmt.Println(Summation(5))
}
