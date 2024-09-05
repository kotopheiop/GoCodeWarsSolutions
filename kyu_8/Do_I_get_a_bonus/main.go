package main

import (
	"fmt"
)

func BonusTime(salary int, bonus bool) string {
	if bonus {
		return fmt.Sprintf("£%d", salary*10)
	}

	return fmt.Sprintf("£%d", salary)
}

func main() {
	fmt.Println(BonusTime(100, false))   // £100
	fmt.Println(BonusTime(55000, false)) // £55000
	fmt.Println(BonusTime(100, true))    // £1000
	fmt.Println(BonusTime(14000, true))  // £140000
}
