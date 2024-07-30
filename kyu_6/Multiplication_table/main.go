package main

import "fmt"

func MultiplicationTable(size int) [][]int {
	table := make([][]int, size)
	for i := range table {
		table[i] = make([]int, size)
		for j := range table[i] {
			table[i][j] = (i + 1) * (j + 1)
		}
	}
	return table
}

func main() {
	fmt.Println(MultiplicationTable(3))
}
