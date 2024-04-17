package main

import "fmt"

func main() {
	chickenMenu := 10.35
	fishMenu := 12.40
	vegeterianMenu := 8.15
	delivery := 2.5

	var chickenMenuCount, fishMenuCount, vegeterianMenuCount float64

	fmt.Scanln(&chickenMenuCount)
	fmt.Scanln(&fishMenuCount)
	fmt.Scanln(&vegeterianMenuCount)

	allCostsWithoutDessert := chickenMenu*chickenMenuCount +
		fishMenu*fishMenuCount + vegeterianMenu*vegeterianMenuCount

	dessertCosts := allCostsWithoutDessert * 0.20

	allMoneyNeeded := allCostsWithoutDessert + dessertCosts + delivery

	fmt.Println(allMoneyNeeded)
}
