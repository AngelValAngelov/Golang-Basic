package main

import "fmt"

func main() {
	var pricePerSquareMeter = 7.61

	var allSquareMeters float64
	fmt.Scanln(&allSquareMeters)

	discount := pricePerSquareMeter * allSquareMeters * 0.18
	allSum := pricePerSquareMeter*allSquareMeters - discount

	fmt.Printf("The final price is: %f lv.", allSum)
	fmt.Printf("The discount is: %f lv.", discount)
}
