package main

import (
	"fmt"
	"math"
)

func main() {
	puzzlePrice := 2.60
	speakingToyPrice := 3.00
	bearToyPrice := 4.10
	minionPrice := 8.20
	toyTruckPrice := 2.00

	var tripPrice, puzzlesCount, speakingToysCount, toyBearsCount, minionsCount, toyTrucksCount float64
	fmt.Scanln(&tripPrice)
	fmt.Scanln(&puzzlesCount)
	fmt.Scanln(&speakingToysCount)
	fmt.Scanln(&toyBearsCount)
	fmt.Scanln(&minionsCount)
	fmt.Scanln(&toyTrucksCount)

	toysCount := puzzlesCount + speakingToysCount + toyBearsCount + minionsCount + toyTrucksCount

	turnover := puzzlesCount*puzzlePrice + speakingToysCount*speakingToyPrice +
		toyBearsCount*bearToyPrice + minionsCount*minionPrice + toyTrucksCount*toyTruckPrice

	if toysCount >= 50 {
		turnover *= 0.75
	}

	turnoverAfterRent := turnover * 0.90
	moneyLeft := math.Abs(turnoverAfterRent - tripPrice)

	if turnoverAfterRent >= tripPrice {

		fmt.Printf("Yes! %.2f lv left.", moneyLeft)
	} else {
		fmt.Printf("Not enough money! %.2f lv needed.", moneyLeft)
	}
}
