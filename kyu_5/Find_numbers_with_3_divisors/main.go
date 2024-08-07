package main

import (
	"fmt"
	"math"
	"sort"
)

// Решение не засчитано
// Ссылка на кату: https://www.codewars.com/kata/65eb2c4c274bd91c27b38d32
// Разобраться с большими числами

func sieveOfEratosthenes(limit int64) []int64 {
	isPrime := make([]bool, limit+1)
	for i := range isPrime {
		isPrime[i] = true
	}
	isPrime[0], isPrime[1] = false, false

	for p := int64(2); p*p <= limit; p++ {
		if isPrime[p] {
			for i := p * p; i <= limit; i += p {
				isPrime[i] = false
			}
		}
	}

	var primes []int64
	for p := int64(2); p <= limit; p++ {
		if isPrime[p] {
			primes = append(primes, p)
		}
	}
	return primes
}

func Solution(n, m int) []int {
	n64, m64 := int64(n), int64(m)
	limit := int64(math.Pow(float64(m64), 1.0/4.0))
	primes := sieveOfEratosthenes(limit)

	var result []int
	for _, p := range primes {
		pFourth := p * p * p * p
		if pFourth >= n64 && pFourth <= m64 {
			result = append(result, int(pFourth))
		}
	}

	if len(result) == 0 {
		return []int{}
	}

	sort.Slice(result, func(i, j int) bool { return result[i] < result[j] })
	return result
}

func main() {
	fmt.Println(Solution(2, 100))        // [16, 81]
	fmt.Println(Solution(10000, 100000)) // [14641, 28561, 83521]
	fmt.Println(Solution(624, 625))      // [625]
	fmt.Println(Solution(625, 626))      // [625]
	fmt.Println(Solution(734, 735))      // []
}
