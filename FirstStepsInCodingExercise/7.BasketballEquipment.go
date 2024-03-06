package main

import "fmt"

func main() {
	var yearTax float64
	fmt.Scanln(&yearTax)

	shoes := yearTax * 0.60
	clothes := shoes * 0.80
	ball := clothes * 0.25
	accessories := ball * 0.20

	allMoneyNeeded := shoes + clothes + ball + accessories + yearTax

	fmt.Println(allMoneyNeeded)
}
