package main

import (
	"fmt"
	"time"
)

func Past(h, m, s int) int {
	midnight, _ := time.Parse(time.TimeOnly, "00:00:00")

	endTime := midnight.
		Add(time.Hour * time.Duration(h)).
		Add(time.Minute * time.Duration(m)).
		Add(time.Second * time.Duration(s))

	duration := endTime.Sub(midnight)
	milliseconds := int(duration.Milliseconds())

	return milliseconds
}

func main() {
	fmt.Println(Past(0, 1, 1)) // 61000
	fmt.Println(Past(1, 1, 1)) // 3661000
	fmt.Println(Past(0, 0, 0)) // 0
	fmt.Println(Past(1, 0, 1)) // 3601000
	fmt.Println(Past(1, 0, 0)) // 3600000
}
