package main

import "fmt"

func main() {
	var firstTime, SecondTime, ThirdTime int

	fmt.Scanln(&firstTime)
	fmt.Scanln(&SecondTime)
	fmt.Scanln(&ThirdTime)

	allTimeInSeconds := firstTime + SecondTime + ThirdTime
	allTimeMinutes := allTimeInSeconds / 60
	allTimeSeconds := allTimeInSeconds % 60

	fmt.Printf("%d:%02d", allTimeMinutes, allTimeSeconds)
}
