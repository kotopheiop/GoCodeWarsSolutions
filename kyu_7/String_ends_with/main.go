package main

import "fmt"

func solution(str, ending string) bool {
	strLen := len(str)
	endingLen := len(ending)

	if strLen < endingLen {
		return false
	}

	if str[strLen-endingLen:] == ending {
		return true
	}

	return false
}

func main() {
	fmt.Println(solution("banana", "ana")) // true
	fmt.Println(solution("abc", "d"))      // false
}
