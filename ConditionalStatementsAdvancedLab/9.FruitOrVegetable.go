package main

import "fmt"

func main() {
	var food string
	fmt.Scanln(&food)

	switch food {
	case "banana", "apple", "kiwi", "cherry", "lemon", "grapes":
		fmt.Println("fruit")
	case "tomato", "cucumber", "pepper", "carrot":
		fmt.Println("vegetable")
	default:
		fmt.Println("unknown")
	}
}
