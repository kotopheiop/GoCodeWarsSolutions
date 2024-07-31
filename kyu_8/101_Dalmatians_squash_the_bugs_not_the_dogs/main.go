package main

import "fmt"

func HowManyDalmatians(dogCount int) string {
	var result string

	switch {
	case dogCount == 101:
		result = "101 DALMATIONS!!!"
	case dogCount <= 10:
		result = "Hardly any"
	case dogCount > 10 && dogCount <= 50:
		result = "More than a handful!"
	case dogCount > 50 && dogCount <= 100:
		result = "Woah that's a lot of dogs!"
	}

	return result
}

func main() {
	fmt.Println(HowManyDalmatians(26))  //More than a handful!
	fmt.Println(HowManyDalmatians(8))   //Hardly any
	fmt.Println(HowManyDalmatians(14))  //More than a handful!
	fmt.Println(HowManyDalmatians(80))  //Woah that's a lot of dogs!
	fmt.Println(HowManyDalmatians(100)) //Woah that's a lot of dogs!
	fmt.Println(HowManyDalmatians(50))  //More than a handful!
	fmt.Println(HowManyDalmatians(10))  //Hardly any
	fmt.Println(HowManyDalmatians(101)) //101 DALMATIONS!!!
}
