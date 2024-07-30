package main

import "fmt"

func FindUniq(arr []float32) (result float32) {
	m := make(map[float32]int)
	for _, number := range arr {
		m[number]++
	}
	for number, count := range m {
		if count == 1 {
			result = number
			break
		}
	}

	return
}

func main() {
	fmt.Println(FindUniq([]float32{1, 1, 1, 2, 1, 1}))
	fmt.Println(FindUniq([]float32{0, 0, 0.55, 0, 0}))
}
