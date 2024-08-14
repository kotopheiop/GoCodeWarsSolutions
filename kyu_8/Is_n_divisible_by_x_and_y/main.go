package main

import "fmt"

func IsDivisible(n, x, y int) bool {
	return n%x == 0 && n%y == 0
}

func main() {
	fmt.Println(IsDivisible(3, 3, 4))    // false
	fmt.Println(IsDivisible(12, 3, 4))   // true
	fmt.Println(IsDivisible(100, 5, 10)) // true
	fmt.Println(IsDivisible(100, 5, 3))  // false
}
