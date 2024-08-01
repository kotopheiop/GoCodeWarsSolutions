package main

import "fmt"

func Persistence(num int) int {
	count := 0
	for num >= 10 {
		num = multiply(num)
		count++
	}

	return count
}

func multiply(num int) int {
	product := 1
	for num > 0 {
		digit := num % 10
		product *= digit
		num /= 10
	}

	return product
}

func main() {
	fmt.Println(Persistence(39))  // 3
	fmt.Println(Persistence(999)) // 4
	fmt.Println(Persistence(4))   // 0
	fmt.Println(Persistence(25))  // 2
}
