package main

import "fmt"

func main() {
	var month string
	var nightCount int
	fmt.Scanln(&month)
	fmt.Scanln(&nightCount)

	studio := 0.0
	apartment := 0.0

	if month == "May" || month == "October" {
		studio = 50.0 * float64(nightCount)
		apartment = 65.0 * float64(nightCount)
	} else if month == "June" || month == "September" {
		studio = 75.20 * float64(nightCount)
		apartment = 68.70 * float64(nightCount)
	} else {
		studio = 76.0 * float64(nightCount)
		apartment = 77.0 * float64(nightCount)
	}

	if (nightCount > 7 && nightCount <= 14) && (month == "May" || month == "October") {
		studio = 50 * float64(nightCount) * 0.95
	} else if nightCount > 14 && (month == "May" || month == "October") {
		studio = 50 * float64(nightCount) * 0.70
	} else if nightCount > 14 && (month == "June" || month == "September") {
		studio = 75.20 * float64(nightCount) * 0.80
	}

	if nightCount > 14 {
		if month == "May" || month == "October" {
			apartment = 65 * float64(nightCount) * 0.90
		} else if month == "June" || month == "September" {
			apartment = 68.70 * float64(nightCount) * 0.90
		} else {
			apartment = 77 * float64(nightCount) * 0.90
		}
	}
	fmt.Printf("Apartment: %.2f lv.\n", apartment)
	fmt.Printf("Studio: %.2f lv.", studio)
}
