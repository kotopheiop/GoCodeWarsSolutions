package main

import (
	"fmt"
	"math"
	"strings"
)

func Code(s string) string {
	n := int(math.Ceil(math.Sqrt(float64(len(s)))))
	for len(s) < n*n {
		s += "\x0B"
	}
	lines := make([]string, n)
	for i := 0; i < n; i++ {
		lines[i] = s[i*n : (i+1)*n]
	}
	result := ""
	for i := 0; i < n; i++ {
		for j := n - 1; j >= 0; j-- {
			if len(lines[j]) > i {
				result += string(lines[j][i])
			}
		}
		if i != n-1 {
			result += "\n"
		}
	}
	return result
}

func Decode(s string) string {
	lines := strings.Split(s, "\n")
	n := len(lines)
	result := ""
	for i := n - 1; i >= 0; i-- {
		for j := 0; j < n; j++ {
			if len(lines[j]) > i {
				result += string(lines[j][i])
			}
		}
	}
	return strings.Replace(result, "\x0B", "", -1)
}

func main() {
	t := "I.was.going.fishing.that.morning.at.ten.o'clock"
	fmt.Println("Оригинальный текст:", t)

	coded := Code(t)

	fmt.Println("Кодированный текст:\n", coded)
	fmt.Println("Декодированный текст:", Decode(coded))
}
