package main

import (
	"fmt"
	"math"
)

func main() {
	var radians float32
	fmt.Scanln(&radians)

	degrees := radians * 180 / math.Pi

	fmt.Println(degrees)
}
