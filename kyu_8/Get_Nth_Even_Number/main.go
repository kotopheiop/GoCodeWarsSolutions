package main

import "fmt"

func NthEven(n int) int {
	return 2 * (n - 1)
}

func main() {
	fmt.Println(NthEven(1))   // 0
	fmt.Println(NthEven(100)) // 198
}
