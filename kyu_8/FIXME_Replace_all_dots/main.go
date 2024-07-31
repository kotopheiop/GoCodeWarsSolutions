package main

import (
	"fmt"
	"strings"
)

func ReplaceDots(str string) string {
	return strings.Replace(str, ".", "-", -1)
}

func main() {
	fmt.Println(ReplaceDots(""))              //""
	fmt.Println(ReplaceDots("no dots"))       //"no dots"
	fmt.Println(ReplaceDots("one.two.three")) //"one-two-three"
	fmt.Println(ReplaceDots("..."))           //"---"
}
