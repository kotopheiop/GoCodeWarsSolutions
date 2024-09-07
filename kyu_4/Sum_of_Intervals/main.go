package main

import (
	"fmt"
	"sort"
)

func SumOfIntervals(intervals [][2]int) int {
	if len(intervals) == 0 {
		return 0
	}

	// Сортируем по начальной точке
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged := [][2]int{intervals[0]}

	for _, interval := range intervals[1:] {
		last := &merged[len(merged)-1]
		if interval[0] <= last[1] { // Пересекающиеся интервалы
			if interval[1] > last[1] {
				last[1] = interval[1] // Обновляем конечную точку
			}
		} else {
			// Добавляем новый интервал
			merged = append(merged, interval)
		}
	}

	// Суммируем длины
	totalLength := 0
	for _, interval := range merged {
		totalLength += interval[1] - interval[0]
	}

	return totalLength
}

func main() {
	fmt.Println(SumOfIntervals([][2]int{{1, 5}}))                                // 4
	fmt.Println(SumOfIntervals([][2]int{{1, 4}, {7, 10}, {3, 5}}))               // 7
	fmt.Println(SumOfIntervals([][2]int{{-1_000_000_000, 1_000_000_000}}))       // 2_000_000_000
	fmt.Println(SumOfIntervals([][2]int{{0, 20}, {-100_000_000, 10}, {30, 40}})) // 100_000_030
}
