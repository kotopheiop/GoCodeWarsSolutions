package main

import (
	"fmt"
	"strconv"
	"strings"
)

func CountBits(n uint) int {
	n_bin := strconv.FormatInt(int64(n), 2)

	return strings.Count(string(n_bin), "1")
}

func main() {
	fmt.Println(CountBits(1234)) // 5
}
