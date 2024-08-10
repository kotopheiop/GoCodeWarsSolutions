package main

import "fmt"

func MakeNegative(x int) int {
	if x <= 0 {
		return x
	}

	return -x
}

func main() {
	fmt.Println(MakeNegative(-2)) // -2
	fmt.Println(MakeNegative(1))  // -1
	fmt.Println(MakeNegative(0))  // 0

}
