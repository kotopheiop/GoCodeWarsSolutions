package main

import "fmt"

func Opposite(value int) int {
	return -value
}

func main() {
	fmt.Println(Opposite(8))  // -8
	fmt.Println(Opposite(-8)) // 8
	fmt.Println(Opposite(0))  // 0

}
