package main

import "fmt"

func main() {
	rose := 5
	dahlia := 3.80
	tulip := 2.80
	narcissus := 3
	gladiolus := 2.50

	var typeFlowers string
	var flowersCount int
	var budget int
	fmt.Scanln(&typeFlowers)
	fmt.Scanln(&flowersCount)
	fmt.Scanln(&budget)
	cost := 0.0

	if typeFlowers == "Roses" && flowersCount > 80 {
		cost = (float64(flowersCount) * float64(rose)) * 0.90
	} else if typeFlowers == "Roses" && flowersCount <= 80 {
		cost = float64(flowersCount) * float64(rose)
	} else if typeFlowers == "Dahlias" && flowersCount > 90 {
		cost = (float64(flowersCount) * float64(dahlia)) * 0.85
	} else if typeFlowers == "Dahlias" && flowersCount <= 90 {
		cost = float64(flowersCount) * float64(dahlia)
	} else if typeFlowers == "Tulips" && flowersCount > 80 {
		cost = (float64(flowersCount) * float64(tulip)) * 0.85
	} else if typeFlowers == "Tulips" && flowersCount <= 80 {
		cost = float64(flowersCount) * float64(tulip)
	} else if typeFlowers == "Narcissus" && flowersCount < 120 {
		cost = (float64(flowersCount) * float64(narcissus)) * 1.15
	} else if typeFlowers == "Narcissus" && flowersCount >= 120 {
		cost = float64(flowersCount) * float64(narcissus)
	} else if typeFlowers == "Gladiolus" && flowersCount < 80 {
		cost = (float64(flowersCount) * float64(gladiolus)) * 1.20
	} else if typeFlowers == "Gladiolus" && flowersCount >= 80 {
		cost = float64(flowersCount) * float64(gladiolus)
	}

	if float64(budget) >= cost {
		moneyLeft := float64(budget) - cost
		fmt.Printf("Hey, you have a great garden with %d %s and %.2f leva left.", flowersCount, typeFlowers, moneyLeft)
	} else {
		moneyNeeded := cost - float64(budget)
		fmt.Printf("Not enough money, you need %.2f leva more.", moneyNeeded)
	}
}
