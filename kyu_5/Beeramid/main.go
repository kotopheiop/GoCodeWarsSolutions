package main

import (
	"fmt"
)

func Beeramid(bonus int, price float64) int {
	beerGlassesBuy := float64(bonus) / price
	var levels int
	var totalBeerGlass int

	for {
		beerGlassAtLevel := levels * levels
		if totalBeerGlass+beerGlassAtLevel > int(beerGlassesBuy) {
			break
		}
		totalBeerGlass += beerGlassAtLevel
		levels += 1
	}

	if levels == 0 {
		return levels
	}

	return levels - 1
}

func main() {
	fmt.Println(Beeramid(5000, 3)) // 16
	fmt.Println(Beeramid(1500, 2)) // 12
	fmt.Println(Beeramid(-1, 4.0)) // 0
}
