package main

import (
	"fmt"
	"math"
)

func main() {
	var budget int
	var season string
	var fishersCount int

	fmt.Scanln(&budget)
	fmt.Scanln(&season)
	fmt.Scanln(&fishersCount)

	boatRent := 0.0

	if season == "Spring" {
		boatRent = 3000
	} else if season == "Summer" || season == "Autumn" {
		boatRent = 4200
	} else {
		boatRent = 2600
	}

	if fishersCount <= 6 {
		boatRent *= 0.90
	} else if fishersCount <= 11 {
		boatRent *= 0.85
	} else {
		boatRent *= 0.75
	}

	if fishersCount%2 == 0 && season != "Autumn" {
		boatRent *= 0.95
	}

	if float64(budget) >= boatRent {
		moneyLeft := float64(budget) - boatRent
		fmt.Printf("Yes! You have %.2f leva left.", moneyLeft)
	} else {
		moneyNeeded := math.Abs(float64(budget) - boatRent)
		fmt.Printf("Not enough money! You need %.2f leva.", moneyNeeded)
	}
}
