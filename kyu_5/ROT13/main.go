package main

import (
	"fmt"
	"strings"
)

func Rot13(msg string) string {
	var result strings.Builder
	for _, char := range msg {
		if char >= 'a' && char <= 'z' {
			result.WriteRune('a' + (char-'a'+13)%26)
		} else if char >= 'A' && char <= 'Z' {
			result.WriteRune('A' + (char-'A'+13)%26)
		} else {
			result.WriteRune(char)
		}
	}
	return result.String()
}

func main() {
	fmt.Println(Rot13("EBG13 rknzcyr."))                    // ROT13 example.
	fmt.Println(Rot13("This is my first ROT13 excercise!")) // Guvf vf zl svefg EBG13 rkprepvfr!
}
