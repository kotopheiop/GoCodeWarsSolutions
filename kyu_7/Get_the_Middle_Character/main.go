package main

import (
	"fmt"
)

func GetMiddle(s string) string {
	length := len(s)

	if length%2 == 0 {
		return s[length/2-1 : length/2+1]
	}

	return string(s[length/2])
}

func main() {
	fmt.Println(GetMiddle("testing")) // t
	fmt.Println(GetMiddle("test"))    // es
	fmt.Println(GetMiddle("middle"))  // dd
	fmt.Println(GetMiddle("A"))       // A
}
