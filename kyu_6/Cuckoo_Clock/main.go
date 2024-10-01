package main

import (
	"fmt"
	"time"
)

func CuckooClock(initialTime string, chimes int) string {
	startTime, _ := time.Parse("15:04", initialTime)

	minutes := startTime.Minute()
	roundedMinutes := ((minutes + 14) / 15) * 15
	if roundedMinutes >= 60 {
		startTime = startTime.Add(time.Hour).Truncate(time.Hour)
	} else {
		startTime = startTime.Truncate(time.Hour).Add(time.Duration(roundedMinutes) * time.Minute)
	}

	// Количество ударов кукушки в данный момент времени
	chimesAtTime := func(t time.Time) int {
		if t.Minute() == 0 {
			hour := t.Hour()
			if hour == 0 || hour == 12 { // Полночь или 12:00 - 12 раз
				return 12
			} else if hour > 12 { // 13:00 -> 1 раз, 14:00 -> 2 раза и т.д.
				return hour - 12
			}
			return hour
		}
		return 1
	}

	currentChimes := chimesAtTime(startTime)
	chimes -= currentChimes

	for chimes > 0 {
		startTime = startTime.Add(15 * time.Minute)
		chimes -= chimesAtTime(startTime)
	}

	return startTime.Format("03:04")
}

func main() {
	fmt.Println(CuckooClock("03:38", 19))  // "06:00"
	fmt.Println(CuckooClock("10:00", 10))  // "10:00"
	fmt.Println(CuckooClock("08:17", 113)) // "08:00"
	fmt.Println(CuckooClock("08:17", 200)) // "05:45"
}
