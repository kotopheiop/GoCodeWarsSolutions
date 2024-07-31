package main

import "fmt"

func Multiply(a, b int) int {
	return a * b
}

func main() {
	fmt.Println(Multiply(1, 1))  // 1
	fmt.Println(Multiply(2, 5))  // 10
	fmt.Println(Multiply(5, 10)) // 50
	fmt.Println(Multiply(5, 0))  // 0
	fmt.Println(Multiply(0, 5))  // 0
}
