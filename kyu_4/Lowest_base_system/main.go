package main

import (
	"fmt"
	"math"
	"math/big"
)

func GetMinBase(n uint64) uint64 {
	// Оценка верхней границы для k на основе логарифма числа n.
	// Это помогает предотвратить бесконечный цикл и ограничивает количество проверок.
	maxK := int(math.Log2(float64(n))) + 2

	for k := 2; k <= maxK; k++ {
		base := findBase(n, k)
		if base != 0 {
			return base
		}
	}

	// Если подходящая база не найдена, возвращаем n-1 (в таком случае n = 11...1 в базе n-1)
	return n - 1
}

func findBase(n uint64, k int) uint64 {
	low := big.NewInt(2)
	high := big.NewInt(0).Exp(big.NewInt(int64(n)), big.NewInt(1), nil)
	high.Div(high, big.NewInt(int64(k))) // Верхняя граница на базе

	// Бинарный поиск по возможным значениям базы
	for low.Cmp(high) <= 0 {
		// Вычисляем середину диапазона
		mid := big.NewInt(0).Add(low, high)
		mid.Div(mid, big.NewInt(2))

		sum := big.NewInt(1)           // Начальная сумма (1)
		curr := big.NewInt(1)          // Текущая степень базы
		base := big.NewInt(0).Set(mid) // Текущая предполагаемая база

		// Считаем сумму последовательности 1 + b + b^2 + ... + b^(k-1)
		for i := 1; i < k; i++ {
			curr.Mul(curr, base)
			sum.Add(sum, curr)
			// Если сумма превысила n, прерываем цикл
			if sum.Cmp(big.NewInt(int64(n))) > 0 {
				break
			}
		}

		// Если сумма равна n, мы нашли нужную базу
		if sum.Cmp(big.NewInt(int64(n))) == 0 {
			return mid.Uint64()
		} else if sum.Cmp(big.NewInt(int64(n))) < 0 {
			low = mid.Add(mid, big.NewInt(1))
		} else {
			high = mid.Sub(mid, big.NewInt(1))
		}
	}

	return 0
}

func main() {
	fmt.Println(GetMinBase(3))             // 2
	fmt.Println(GetMinBase(7))             // 2
	fmt.Println(GetMinBase(21))            // 4
	fmt.Println(GetMinBase(57))            // 7
	fmt.Println(GetMinBase(1111))          // 10
	fmt.Println(GetMinBase(1000002))       // 1000001
	fmt.Println(GetMinBase(1000000002))    // 1000000001
	fmt.Println(GetMinBase(1000000000000)) // 999999999999
	fmt.Println(GetMinBase(1001001))       // 1000
	fmt.Println(GetMinBase(1001001001))    // 1000
	fmt.Println(GetMinBase(1001001001001)) // 1000
	fmt.Println(GetMinBase(2500050001))    // 50000
}
