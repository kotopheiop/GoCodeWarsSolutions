package main

import (
	"fmt"
	"strconv"
)

func FakeBin(x string) string {
	var res string
	if len(x) == 0 {
		return res
	}

	for _, val := range x {
		num, _ := strconv.Atoi(string(val))
		if num >= 5 {
			res += "1"
		} else {
			res += "0"
		}
	}

	return res
}

func main() {
	fmt.Println(FakeBin("45385593107843568")) // 01011110001100111
}
