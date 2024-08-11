package main

import (
	"fmt"
	"strings"
)

func StringToArray(str string) []string {
	return strings.Fields(str)
}

func main() {
	fmt.Println(StringToArray("I love arrays they are my favorite")) // [I love arrays they are my favorite]
}
