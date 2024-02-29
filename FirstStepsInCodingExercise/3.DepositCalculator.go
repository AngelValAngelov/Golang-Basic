package main

import "fmt"

func main() {
	var depositSum float32
	fmt.Scanln(&depositSum)

	var depositTerm int
	fmt.Scanln(&depositTerm)

	var interestPersent float32
	fmt.Scanln(&interestPersent)

	allSum := depositSum + float32(depositTerm)*((depositSum*interestPersent/100)/12)

	fmt.Println(allSum)
}
    