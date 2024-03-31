package main

import "fmt"

func main() {
	var drinkType, city string
	var quantity float32

	fmt.Scanln(&drinkType)
	fmt.Scanln(&city)
	fmt.Scanln(&quantity)

	if city == "Sofia" {
		if drinkType == "coffee" {
			result := quantity * 0.50
			fmt.Println(result)
		} else if drinkType == "water" {
			result := quantity * 0.80
			fmt.Println(result)
		} else if drinkType == "beer" {
			result := quantity * 1.20
			fmt.Println(result)
		} else if drinkType == "sweets" {
			result := quantity * 1.45
			fmt.Println(result)
		} else if drinkType == "peanuts" {
			result := quantity * 1.60
			fmt.Println(result)
		}
	} else if city == "Plovdiv" {
		if drinkType == "coffee" {
			result := quantity * 0.40
			fmt.Println(result)
		} else if drinkType == "water" {
			result := quantity * 0.70
			fmt.Println(result)
		} else if drinkType == "beer" {
			result := quantity * 1.15
			fmt.Println(result)
		} else if drinkType == "sweets" {
			result := quantity * 1.30
			fmt.Println(result)
		} else if drinkType == "peanuts" {
			result := quantity * 1.50
			fmt.Println(result)
		}
	} else if city == "Varna" {
		if drinkType == "coffee" {
			result := quantity * 0.45
			fmt.Println(result)
		} else if drinkType == "water" {
			result := quantity * 0.70
			fmt.Println(result)
		} else if drinkType == "beer" {
			result := quantity * 1.10
			fmt.Println(result)
		} else if drinkType == "sweets" {
			result := quantity * 1.35
			fmt.Println(result)
		} else if drinkType == "peanuts" {
			result := quantity * 1.55
			fmt.Println(result)
		}
	}
}
