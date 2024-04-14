package main

import "fmt"

func main() {
	var budget float64
	var typeSeason string

	fmt.Scanln(&budget)
	fmt.Scanln(&typeSeason)   

	destination := ""
	typeHousing := ""
	price := 0.0

	if budget <= 100 {
		destination = "Bulgaria"
		if typeSeason == "summer" {
			typeHousing = "Camp"
			price = budget * 0.30
		} else {
			typeHousing = "Hotel"
			price = budget * 0.70
		}
	} else if budget <= 1000 {
		destination = "Balkans"
		if typeSeason == "summer" {
			typeHousing = "Camp"
			price = budget * 0.40
		} else {
			typeHousing = "Hotel"
			price = budget * 0.80
		}
	} else {
		typeHousing = "Hotel"
		destination = "Europe"
		price = budget * 0.90
	}

	fmt.Printf("Somewhere in %s\n", destination)
	fmt.Printf("%s - %.2f", typeHousing, price)
}
