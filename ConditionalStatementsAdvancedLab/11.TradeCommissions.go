package main

import "fmt"

func main() {
	var city string
	var money float32

	fmt.Scanln(&city)
	fmt.Scanln(&money)


	if city == "Sofia" && (money >= 0 && money <= 500) {
		fmt.Printf("%.2f", (money * 0.05))
	} else if city == "Sofia" && (money >= 500 && money <= 1000) {
		fmt.Printf("%.2f", (money * 0.07))
	} else if city == "Sofia" && (money >= 1000 && money <= 10000) {
		fmt.Printf("%.2f", (money * 0.08))
	} else if city == "Sofia" && (money >= 10000) {
		fmt.Printf("%.2f", (money * 0.12))
	} else if city == "Varna" && money >= 0 && money <= 500 {
		fmt.Printf("%.2f", (money * 0.045))
	} else if city == "Varna" && money >= 500 && money <= 1000 {
		fmt.Printf("%.2f", (money * 0.075))
	} else if city == "Varna" && money >= 1000 && money <= 10000 {
		fmt.Printf("%.2f", (money * 0.10))
	} else if city == "Varna" && money >= 10000 {
		fmt.Printf("%.2f", (money * 0.13))
	} else if city == "Plovdiv" && money >= 0 && money <= 500 {
		fmt.Printf("%.2f", (money * 0.055))
	} else if city == "Plovdiv" && money >= 500 && money <= 1000 {
		fmt.Printf("%.2f", (money * 0.08))
	} else if city == "Plovdiv" && money >= 1000 && money <= 10000 {
		fmt.Printf("%.2f", (money * 0.12))
	} else if city == "Plovdiv" && money >= 10000 {
		fmt.Printf("%.2f", (money * 0.145))
	} else {
		fmt.Println("error")
	}
}
