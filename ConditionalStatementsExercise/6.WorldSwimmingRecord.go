package main

import (
	"fmt"
	"math"
)

func main() {
	var recordTime, distance, timeSecondsPerMeter float64
	fmt.Scanln(&recordTime)
	fmt.Scanln(&distance)
	fmt.Scanln(&timeSecondsPerMeter)

	swimmingTime := distance * timeSecondsPerMeter
	timeTimeLoosing := math.Floor(distance/15) * 12.5
	allTime := swimmingTime + timeTimeLoosing

	if recordTime <= allTime {
		secondsNeeded := allTime - recordTime
		fmt.Printf("No, he failed! He was %.2f seconds slower.", secondsNeeded)
	} else {
		fmt.Printf("Yes, he succeeded! The new world record is %.2f seconds.", allTime)
	}
}
