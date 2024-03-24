package main

import (
	"fmt"
	"math"	
)

func main() {

	var tvseries string
	var lengthEpisode, lengthBreak int


	fmt.Scanln(&tvseries)
	fmt.Scanln(&lengthEpisode)
	fmt.Scanln(&lengthBreak)

	lunch := float64(lengthBreak) / 8
	timeBreak := float64(lengthBreak) / 4

	leftOver := float64(lengthBreak) - lunch - float64(timeBreak)

	result := math.Ceil(leftOver - float64(lengthEpisode))
	needed := math.Ceil(float64(lengthEpisode) - leftOver)


	if leftOver >= float64(lengthEpisode) {
		fmt.Printf("You have enough time to watch %s and left with %.0f minutes free time.", tvseries, result)
	} else {
		fmt.Printf("You don't have enough time to watch %s, you need %.0f more minutes.", tvseries, needed)
	}
}
