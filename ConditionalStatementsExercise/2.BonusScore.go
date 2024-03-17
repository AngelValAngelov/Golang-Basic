package main

import "fmt"

func main() {
	var number int
	fmt.Scanln(&number)

	bonus := 0.0
	result := 0.0

	if number < 100 {
		bonus += 5
	} else if number <= 1000 {
		bonus += float64(number) * 0.20
	} else if number > 1000 {
		bonus += float64(number) * 0.10
	}

	if number%2 == 0 {
		bonus += 1
	} else if number % 5 == 0 {
		bonus += 2
	}

	result = float64(number) + bonus

	fmt.Println(bonus)
	fmt.Println(result)
}
