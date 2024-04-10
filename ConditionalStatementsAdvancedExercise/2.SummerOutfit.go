package main

import "fmt"

func main() {
	var degrees int
	var time string

	fmt.Scanln(&degrees)
	fmt.Scanln(&time)

	if time == "Morning" && degrees >= 10 && degrees <= 18 {
		fmt.Printf("It's %d degrees, get your Sweatshirt and Sneakers.", degrees)
	} else if time == "Morning" && degrees > 18 && degrees <= 24 {
		fmt.Printf("It's %d degrees, get your Shirt and Moccasins.", degrees)
	} else if time == "Morning" && degrees >= 25 {
		fmt.Printf("It's %d degrees, get your T-Shirt and Sandals.", degrees)
	} else if time == "Afternoon" && degrees >= 10 && degrees <= 18 {
		fmt.Printf("It's %d degrees, get your Shirt and Moccasins.", degrees)
	} else if time == "Afternoon" && degrees > 18 && degrees <= 24 {
		fmt.Printf("It's %d degrees, get your T-Shirt and Sandals.", degrees)
	} else if time == "Afternoon" && degrees >= 25 {
		fmt.Printf("It's %d degrees, get your Swim Suit and Barefoot.", degrees)
	} else if time == "Evening" && degrees >= 10 && degrees <= 18 {
		fmt.Printf("It's %d degrees, get your Shirt and Moccasins.", degrees)
	} else if time == "Evening" && degrees > 18 && degrees <= 24 {
		fmt.Printf("It's %d degrees, get your Shirt and Moccasins.", degrees)
	} else if time == "Evening" && degrees >= 25 {
		fmt.Printf("It's %d degrees, get your Shirt and Moccasins.", degrees)
	}
}
