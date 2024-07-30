package main

import "fmt"

func EvenOrOdd(number int) string {
	if number%2 == 0 {
		return "Even"
	} else {
		return "Odd"
	}
}

func main() {
	fmt.Println(EvenOrOdd(12))
	fmt.Println(EvenOrOdd(13))
}
