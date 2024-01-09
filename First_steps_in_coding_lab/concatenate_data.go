package main

import "fmt"

func main() {
	var firstName string
	fmt.Scanln(&firstName)

	var lastName string
	fmt.Scanln(&lastName)

	var years int
	fmt.Scanln(&years)

	var city string
	fmt.Scanln(&city)

	fmt.Printf("You are %s %s, a %d-years old person from %s.", firstName, lastName, years, city)
}
