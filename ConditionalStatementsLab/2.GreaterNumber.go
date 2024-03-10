package main

import "fmt"

func main() {
	var number1, number2 int
	fmt.Scanln(&number1)
	fmt.Scanln(&number2)

	if number1 > number2 {
		fmt.Println(number1)
	} else {
		fmt.Println(number2)
	}
}
