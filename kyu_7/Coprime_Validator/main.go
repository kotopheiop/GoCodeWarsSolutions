package main

import "fmt"

func AreCoprime(n, m int) bool {
	for m != 0 {
		n, m = m, n%m
	}
	if n > 1 {
		return false
	}

	return true
}

func main() {
	fmt.Println(AreCoprime(20, 27))
	fmt.Println(AreCoprime(12, 39))
}
