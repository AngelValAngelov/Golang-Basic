package main

import "fmt"

func main() {
	USDtoBGN := 1.79549

	var USDCount float64
	fmt.Scanln(&USDCount)

	allSum := USDtoBGN * USDCount

	fmt.Println(allSum)
}
