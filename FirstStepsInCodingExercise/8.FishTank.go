package main

import "fmt"

func main() {
	var length, width, height int
	var percent float64

	fmt.Scanln(&length)
	fmt.Scanln(&width)
	fmt.Scanln(&height)
	fmt.Scanln(&percent)

	capacity := float64(length*width*height) * 0.001
	litersWaterNeeded := capacity * (1 - (percent / 100))

	fmt.Println(litersWaterNeeded)
}
