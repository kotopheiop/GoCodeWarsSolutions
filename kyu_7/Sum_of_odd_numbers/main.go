package main

import "fmt"

// n * n * n 😁
func RowSumOddNumbers(n int) int {
	if n < 1 {
		return 0
	}

	firstNum := n*n - (n - 1)

	result := 0
	for i := 0; i < n; i++ {
		result += firstNum + 2*i
	}

	return result
}

func main() {
	fmt.Println(RowSumOddNumbers(1))      // 1
	fmt.Println(RowSumOddNumbers(2))      // 8
	fmt.Println(RowSumOddNumbers(100500)) // 1015075125000000
}
