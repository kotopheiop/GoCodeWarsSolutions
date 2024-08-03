package main

import "fmt"

type Tuple struct {
	Char  rune
	Count int
}

func OrderedCount(text string) []Tuple {
	// Если входная строка пустая, возвращаем пустой срез
	if text == "" {
		return []Tuple{}
	}

	var arr []Tuple
	for _, char := range text {
		var index = -1
		for i, value := range arr {
			if char == value.Char {
				index = i
			}
		}

		if index != -1 {
			arr[index].Count++
		} else {
			arr = append(arr, Tuple{char, 1})
		}
	}

	return arr
}

func main() {
	fmt.Println(OrderedCount("abracadabra")) // []Tuple{Tuple{'a', 5}, Tuple{'b', 2}, Tuple{'r', 2}, Tuple{'c', 1}, Tuple{'d', 1}}
	fmt.Println(OrderedCount(""))            // []Tuple{}
}
