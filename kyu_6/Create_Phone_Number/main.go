package main

import "fmt"

func CreatePhoneNumber(numbers [10]uint) string {
	first := numbers[0:3]
	second := numbers[3:6]
	third := numbers[6:]

	phoneString := fmt.Sprintf("(%d%d%d) %d%d%d-%d%d%d%d", first[0], first[1], first[2], second[0], second[1], second[2], third[0], third[1], third[2], third[3])

	return phoneString
}

func main() {
	fmt.Println(CreatePhoneNumber([10]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}))
}
