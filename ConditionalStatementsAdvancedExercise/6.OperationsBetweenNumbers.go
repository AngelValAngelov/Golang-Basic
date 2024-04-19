package main

import (
	"fmt"
)

func main() {
	var num1, num2 int
	var sign string
	fmt.Scan(&num1, &num2, &sign)

	evenOrOdd := "odd"
	if sign == "+" || sign == "-" || sign == "*" {
		result := 0
		if sign == "+" {
			result = num1 + num2
		} else if sign == "-" {
			result = num1 - num2
		} else if sign == "*" {
			result = num1 * num2
		}
		if result%2 == 0 {
			evenOrOdd = "even"
		}
		fmt.Printf("%d %s %d = %d - %s", num1, sign, num2, result, evenOrOdd)
	} else if sign == "/" {
		if num2 == 0 {
			fmt.Printf("Cannot divide %d by zero", num1)
		} else {
			result := float64(num1) / float64(num2)
			fmt.Printf("%d / %d = %.2f", num1, num2, result)
		}
	} else if sign == "%" {
		if num2 == 0 {
			fmt.Printf("Cannot divide %d by zero", num1)
		} else {
			result := num1 % num2
			fmt.Printf("%d %% %d = %d", num1, num2, result)
		}
	}
}
