package main

import "fmt"

func DigitalRoot(n int) int {
	result := 0
	for n >= 10 {
		n = sum(n) - 1
		result = n
	}

	return result
}

func sum(num int) int {
	product := 1
	for num > 0 {
		digit := num % 10
		product += digit
		num /= 10
	}

	return product
}

func main() {
	fmt.Println(DigitalRoot(16))     // 7
	fmt.Println(DigitalRoot(942))    // 6
	fmt.Println(DigitalRoot(132189)) // 6
	fmt.Println(DigitalRoot(493193)) // 2
	fmt.Println(DigitalRoot(0))      // 0

}
