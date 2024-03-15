package main

import (
	"fmt"
	"math"
)

func main() {
	var figure string
	fmt.Scanln(&figure)

	if figure == "square" {
		var sideA float64
		fmt.Scanln(&sideA)

		result := sideA * sideA

		fmt.Printf("%.3f", result)
	} else if figure == "rectangle" {
		var sideA, sideB float64
		fmt.Scanln(&sideA)
		fmt.Scanln(&sideB)

		result := sideA * sideB

		fmt.Printf("%.3f", result)
	} else if figure == "circle" {
		var r float64
		fmt.Scanln(&r)

		result := math.Pi * r * r

		fmt.Printf("%.3f", result)
	} else if figure == "triangle" {
		var sideA, hight float64
		fmt.Scanln(&sideA)
		fmt.Scanln(&hight)

		result := (sideA * hight) / 2

		fmt.Printf("%.3f", result)
	}
}
