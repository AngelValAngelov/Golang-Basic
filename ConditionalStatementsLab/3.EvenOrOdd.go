package main

import "fmt"

func main() {
	var number1 int

	fmt.Scanln(&number1)

	if number1%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}
}
