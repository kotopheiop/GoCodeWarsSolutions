package main

import (
	"fmt"
	"strings"
)

func MakeUpperCase(str string) string {
	return strings.ToUpper(str)
}

func main() {
	fmt.Println(MakeUpperCase("Hello World"))        // HELLO WORLD
	fmt.Println(MakeUpperCase("1,2,3 hello world!")) // 1,2,3 HELLO WORLD!
}
