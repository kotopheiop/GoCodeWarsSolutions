package main

import (
	"fmt"
	"regexp"
)

func GetCount(str string) (count int) {
	pattern := `[aeiouAEIOU]`

	re, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("Ошибка компиляции регулярного выражения:", err)
		return
	}

	matches := re.FindAllString(str, -1)
	count = len(matches)

	return
}

func main() {
	fmt.Println(GetCount("abraca dabra")) // 5
	fmt.Println(GetCount("abracadabra"))  // 5
	fmt.Println(GetCount("o q"))          // 1
}
