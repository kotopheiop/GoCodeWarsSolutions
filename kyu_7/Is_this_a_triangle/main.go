package main

import "fmt"

func IsTriangle(a, b, c int) bool {
	flag := ""
	if a+b <= c {
		flag = "c"
	} else if a+c <= b {
		flag = "b"
	} else if b+c <= a {
		flag = "a"
	}

	if len(flag) > 0 {
		return false
	}

	return true
}

func main() {
	fmt.Println(IsTriangle(1, 2, 2))
	fmt.Println(IsTriangle(1, 2, 3))
}
