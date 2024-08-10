package main

import (
	"fmt"
	"strings"
)

func FirstNonRepeating(str string) string {
	var result string

	if len(str) == 0 {
		return result
	}

	ords := strings.Split(strings.ToLower(str), "")
	ordsMap := make(map[string]int)
	for _, ord := range ords {
		val, ok := ordsMap[ord]
		if ok {
			ordsMap[ord] = val + 1
		} else {
			ordsMap[ord] = 1
		}
	}

	for _, ord := range ords {
		if ordsMap[ord] == 1 {
			result = ord
			break
		}
	}

	if !strings.Contains(str, result) {
		return strings.ToUpper(result)
	}

	return result
}

func main() {
	fmt.Println(FirstNonRepeating(""))                                     // ""
	fmt.Println(FirstNonRepeating("stress"))                               // t
	fmt.Println(FirstNonRepeating("moonmen"))                              // e
	fmt.Println(FirstNonRepeating("sTreSS"))                               // T
	fmt.Println(FirstNonRepeating("~><#~><"))                              // #
	fmt.Println(FirstNonRepeating("Go hang a salami, I'm a lasagna hog!")) // ,
}
