package main

import "fmt"

func main() {
	var number float32
	fmt.Scanln(&number)

	if -100 <= number && number <= 100 && number != 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
