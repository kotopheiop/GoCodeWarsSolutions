package main

import (
	"fmt"
	"strconv"
)

func NextSmaller(n int) int {
	s := strconv.Itoa(n)
	digits := []rune(s)

	// Шаг 1: Найти самую правую цифру, которая больше цифры справа от неё
	i := len(digits) - 2
	for i >= 0 && digits[i] <= digits[i+1] {
		i--
	}

	// Если не найдена такая цифра, то возвращаем -1
	if i < 0 {
		return -1
	}

	// Шаг 2: Найти наибольшую цифру справа от позиции i, которая меньше чем digits[i]
	j := len(digits) - 1
	for digits[j] >= digits[i] {
		j--
	}

	// Шаг 3: Поменять найденную цифру с digits[i]
	digits[i], digits[j] = digits[j], digits[i]

	// Шаг 4: Перевернуть цифры после позиции i
	reverse(digits[i+1:])

	// Проверка на наличие ведущих нулей
	if digits[0] == '0' {
		return -1
	}

	nextNum, _ := strconv.Atoi(string(digits))
	return nextNum
}

func reverse(runes []rune) {
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
}

func main() {
	fmt.Println(NextSmaller(21))   // 12
	fmt.Println(NextSmaller(531))  // 513
	fmt.Println(NextSmaller(2071)) // 2017

	fmt.Println(NextSmaller(10))  // -1
	fmt.Println(NextSmaller(111)) // -1
	fmt.Println(NextSmaller(100)) // -1
}
