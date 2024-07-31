package main

import "fmt"

func CorrectTail(body string, tail rune) bool {
	lenBody := len(body)
	if rune(body[lenBody-1]) == tail {
		return true
	}

	return false
}

func main() {
	fmt.Println(CorrectTail("Fox", 'x'))     // true
	fmt.Println(CorrectTail("Rhino", 'o'))   // true
	fmt.Println(CorrectTail("Meerkat", 't')) // true
	fmt.Println(CorrectTail("Dog", 't'))     // false
}
