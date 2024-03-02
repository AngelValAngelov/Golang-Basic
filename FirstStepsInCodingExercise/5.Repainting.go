package main

import "fmt"

func main() {
	protectorPerSqM := 1.50
	paintPerliter := 14.50
	paintThinnerPerLiter := 5.00
	bags := 0.40

	var protectorPerSqMCount, paintPerliterCount, paintThinnerPerLiterCount, workHours float64

	fmt.Scanln(&protectorPerSqMCount)
	fmt.Scanln(&paintPerliterCount)
	fmt.Scanln(&paintThinnerPerLiterCount)
	fmt.Scanln(&workHours)

	allProtectorPrice := protectorPerSqM * (protectorPerSqMCount + 2)
	allPaintPrice := paintPerliter * (paintPerliterCount * 1.10)
	allThinnerPrice := paintThinnerPerLiter * paintThinnerPerLiterCount
	mastersCostsPerHour := ((allProtectorPrice + allPaintPrice + allThinnerPrice + bags) * 0.30) * workHours
	allCosts := allProtectorPrice + allPaintPrice + allThinnerPrice + bags + mastersCostsPerHour

	fmt.Println(allCosts)
}
