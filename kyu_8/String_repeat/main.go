package main

import (
	"fmt"
	"strings"
)

func RepeatStr(repetitions int, value string) string {
	return strings.Repeat(value, repetitions)
}

func main() {
	fmt.Println(RepeatStr(4, "a"))      // aaaa
	fmt.Println(RepeatStr(3, "hello ")) // hello hello hello
}
