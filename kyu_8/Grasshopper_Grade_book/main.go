package main

import "fmt"

func GetGrade(a, b, c int) rune {
	avg := (a + b + c) / 3

	if avg >= 90 {
		return 'A'
	} else if avg >= 80 {
		return 'B'
	} else if avg >= 70 {
		return 'C'
	} else if avg >= 60 {
		return 'D'
	} else {
		return 'F'
	}
}

func main() {
	fmt.Println(string(GetGrade(95, 90, 93)))  // A
	fmt.Println(string(GetGrade(70, 70, 100))) // B
	fmt.Println(string(GetGrade(75, 70, 79)))  // C
	fmt.Println(string(GetGrade(58, 62, 70)))  // D
	fmt.Println(string(GetGrade(58, 59, 60)))  // F
}
