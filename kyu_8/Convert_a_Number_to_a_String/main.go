package main

import (
	"fmt"
	"strconv"
)

var NumberToString = strconv.Itoa

func main() {
	fmt.Println(NumberToString(10))    // "10"
	fmt.Println(NumberToString(-10))   // "-10"
	fmt.Println(NumberToString(79585)) // "79585"
}
