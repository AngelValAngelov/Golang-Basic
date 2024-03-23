package main

import (
	"fmt"
	"math"
)

func main() {
	// videocard := 250

	var budget float64
	var videocardsCount, processorCount, ramMemoryCount int

	fmt.Scanln(&budget)
	fmt.Scanln(&videocardsCount)
	fmt.Scanln(&processorCount)
	fmt.Scanln(&ramMemoryCount)

	videoPrice := videocardsCount * 250
	processorsPrice := float64(processorCount) * float64(videoPrice) * 0.35
	ramPrice := float64(ramMemoryCount) * float64(videoPrice) * 0.10
	totalPrice := float64(videoPrice) + processorsPrice + ramPrice
	
	if videocardsCount > processorCount {
    	totalPrice *= 0.85
	}

	result := math.Abs(float64(budget) - totalPrice)

	if budget >= totalPrice {
		
    	fmt.Printf("You have %.2f leva left!", result)
	} else {
    	fmt.Printf("Not enough money! You need %.2f leva more!", result)
	}
}
