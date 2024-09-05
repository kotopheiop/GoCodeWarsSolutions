package main

import (
	"fmt"
	"strconv"
	"strings"
)

func HighAndLow(in string) string {
	inStrSlice := strings.Split(in, " ")
	minNum, _ := strconv.Atoi(inStrSlice[0])
	maxNum := minNum

	for _, v := range inStrSlice[1:] {
		num, _ := strconv.Atoi(v)
		if num > maxNum {
			maxNum = num
		}
		if num < minNum {
			minNum = num
		}
	}

	return strconv.Itoa(maxNum) + " " + strconv.Itoa(minNum)
}

func main() {
	fmt.Println(HighAndLow("1 2 3 4 5"))  // 5 1
	fmt.Println(HighAndLow("1 9 3 4 -5")) // 9 -5
}
