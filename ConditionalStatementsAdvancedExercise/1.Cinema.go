package main

import "fmt"

func main() {
	var kind string
	var rows, cols int

	fmt.Scanln(&kind)
	fmt.Scanln(&rows)
	fmt.Scanln(&cols)

	result := 0.0

	if kind == "Premiere" {
		result = float64(rows) * float64(cols) * 12.00
	} else if kind == "Normal" {
		result = float64(rows) * float64(cols) * 7.5
	} else if kind == "Discount" {
		result = float64(rows) * float64(cols) * 5.00
	}

	fmt.Printf("%.2f leva", result)
}
