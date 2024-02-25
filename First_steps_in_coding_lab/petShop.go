package main

import "fmt"

func main() {
	var dogFoodPackage = 2.5
	var catFoodPackage = 4.0

	var dogFoodPackageCount int
	fmt.Scanln(&dogFoodPackageCount)

	var catFoodPackageCount int
	fmt.Scanln(&catFoodPackageCount)

	allSum := catFoodPackage*float64(catFoodPackageCount) + dogFoodPackage*float64(dogFoodPackageCount)

	fmt.Printf("%f lv.", allSum)
}
