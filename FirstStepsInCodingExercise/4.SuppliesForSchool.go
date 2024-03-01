package main

import "fmt"

func main() {
	pansPack := 5.80
	markerPack := 7.20
	cleaningLiquidPerLiter := 1.20

	var pansPackCount, markerPackCount, cleaningLiquidPerLiterCount, discount float64

	fmt.Scanln(&pansPackCount)
	fmt.Scanln(&markerPackCount)
	fmt.Scanln(&cleaningLiquidPerLiterCount)
	fmt.Scanln(&discount)

	moneyPans := pansPack * pansPackCount
	moneyMarkers := markerPack * markerPackCount
	moneyCleaningStuf := cleaningLiquidPerLiter * cleaningLiquidPerLiterCount
	allMoneyWithoutDiscount := moneyPans + moneyMarkers + moneyCleaningStuf
	finalCost := allMoneyWithoutDiscount - (allMoneyWithoutDiscount * discount / 100)

	fmt.Println(finalCost)
}
