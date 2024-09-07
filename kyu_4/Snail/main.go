package main

import "fmt"

func Snail(snaipMap [][]int) []int {
	if len(snaipMap) == 0 || len(snaipMap[0]) == 0 {
		return []int{}
	}

	result := []int{}
	topRow, bottomRow, leftCol, rightCol := 0, len(snaipMap)-1, 0, len(snaipMap[0])-1

	for topRow <= bottomRow && leftCol <= rightCol {
		for col := leftCol; col <= rightCol; col++ {
			result = append(result, snaipMap[topRow][col])
		}
		topRow++

		for row := topRow; row <= bottomRow; row++ {
			result = append(result, snaipMap[row][rightCol])
		}
		rightCol--

		if topRow <= bottomRow {
			for col := rightCol; col >= leftCol; col-- {
				result = append(result, snaipMap[bottomRow][col])
			}
			bottomRow--
		}

		if leftCol <= rightCol {
			for row := bottomRow; row >= topRow; row-- {
				result = append(result, snaipMap[row][leftCol])
			}
			leftCol++
		}
	}

	return result
}

func main() {
	array := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	result := Snail(array)
	fmt.Println(result) // Output: [1 2 3 6 9 8 7 4 5]

}
