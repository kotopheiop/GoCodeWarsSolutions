package main

import (
	"fmt"
	"strconv"
)

func NextBigger(n int) int {
	s := strconv.Itoa(n)
	digits := []rune(s)

	// Шаг 1: Найти самую правую цифру, которая меньше цифры справа от неё
	i := len(digits) - 2
	for i >= 0 && digits[i] >= digits[i+1] {
		i--
	}

	// Если не найдена такая цифра, то возвращаем -1
	if i < 0 {
		return -1
	}

	// Шаг 2: Найти наименьшую цифру справа от позиции i, которая больше чем digits[i]
	j := len(digits) - 1
	for digits[j] <= digits[i] {
		j--
	}

	// Шаг 3: Поменять найденную цифру с digits[i]
	digits[i], digits[j] = digits[j], digits[i]

	// Шаг 4: Перевернуть цифры после позиции i
	reverse(digits[i+1:])

	nextNum, _ := strconv.Atoi(string(digits))
	return nextNum
}

func reverse(runes []rune) {
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
}

func main() {
	fmt.Println(NextBigger(12))   // 21
	fmt.Println(NextBigger(513))  // 531
	fmt.Println(NextBigger(2017)) // 2071

	fmt.Println(NextBigger(9))   // -1
	fmt.Println(NextBigger(111)) // -1
	fmt.Println(NextBigger(531)) // -1
}
