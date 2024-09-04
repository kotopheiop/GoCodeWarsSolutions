package main

import (
	"fmt"
	"math/big"
)

func IsPrime(n int) bool {
	return big.NewInt(int64(n)).ProbablyPrime(0)
}

func main() {
	fmt.Println(IsPrime(1))
	fmt.Println(IsPrime(2))
	fmt.Println(IsPrime(-1))
}
