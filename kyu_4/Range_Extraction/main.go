package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Solution(list []int) string {
	sort.Ints(list)

	if len(list) == 0 {
		return ""
	}

	var intervals []string
	start := list[0]

	for i := 1; i < len(list); i++ {
		// Проверяем, если текущее число не является непосредственным последователем предыдущего
		if list[i] != list[i-1]+1 {
			if start == list[i-1] {
				// Если интервал состоит из одного числа
				intervals = append(intervals, strconv.Itoa(start))
			} else if list[i-1] == start+1 {
				// Если есть два числа, и разрыв между ними один
				intervals = append(intervals, fmt.Sprintf("%d,%d", start, list[i-1]))
			} else {
				// Иначе, интервал в формате start-end
				intervals = append(intervals, fmt.Sprintf("%d-%d", start, list[i-1]))
			}
			// Новый интервал
			start = list[i]
		}
	}

	// последний интервал
	if start == list[len(list)-1] {
		intervals = append(intervals, strconv.Itoa(start))
	} else if list[len(list)-1] == start+1 {
		intervals = append(intervals, fmt.Sprintf("%d,%d", start, list[len(list)-1]))
	} else {
		intervals = append(intervals, fmt.Sprintf("%d-%d", start, list[len(list)-1]))
	}

	return strings.Join(intervals, ",")
}

func main() {
	fmt.Println(Solution([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}))                                   // 1-15
	fmt.Println(Solution([]int{-10, -9, -8, -6, -3, -2, -1, 0, 1, 3, 4, 5, 7, 8, 9, 10, 11, 14, 15, 17, 18, 19, 20})) // -10--8,-6,-3-1,3-5,7-11,14,15,17-20
	fmt.Println(Solution([]int{-6, -3, -2, -1, 0, 1, 3, 4, 5, 7, 8, 9, 10, 11, 14, 15, 17, 18, 19, 20}))              // -6,-3-1,3-5,7-11,14,15,17-20
}
