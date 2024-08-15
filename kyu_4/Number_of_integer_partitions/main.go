package main

import "fmt"

func Partitions(n int) int {
	memo := make(map[string]int)
	return partitionHelper(n, n, memo)
}

func partitionHelper(n, max int, memo map[string]int) int {
	if n == 0 {
		return 1
	}

	// Если max <= 0 или n < 0, разбиение невозможно
	if max <= 0 || n < 0 {
		return 0
	}

	key := fmt.Sprintf("%d-%d", n, max)
	if val, ok := memo[key]; ok {
		return val
	}

	result := partitionHelper(n-max, max, memo) + partitionHelper(n, max-1, memo)

	memo[key] = result

	return result
}

func main() {
	fmt.Println(Partitions(5))  // 7
	fmt.Println(Partitions(10)) // 42
	fmt.Println(Partitions(25)) // 1958
}
