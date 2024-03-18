package main

import "fmt"

func main() {
	var hours, minutes int
	fmt.Scanln(&hours)
	fmt.Scanln(&minutes)

	timeInSecondsPlusFifteenMinutes := (hours * 60) + minutes + 15

	hoursAsResolt := timeInSecondsPlusFifteenMinutes / 60
	minutesAsResolt := timeInSecondsPlusFifteenMinutes % 60

	if hoursAsResolt > 23 {
		hoursAsResolt -= 24
	}

	fmt.Printf("%d:%02d", hoursAsResolt, minutesAsResolt)
}
