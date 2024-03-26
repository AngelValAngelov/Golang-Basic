package main

import "fmt"

func main() {
	var dayOfWeek string
	fmt.Scanln(&dayOfWeek)

	switch dayOfWeek {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("Working day")
	case "Saturday", "Sunday":
		fmt.Println("Weekend")
	default:
		fmt.Println("Error")
	}
}
