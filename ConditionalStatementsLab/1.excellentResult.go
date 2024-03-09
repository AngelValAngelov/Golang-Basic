package main

import "fmt"

func main() {
	var mark int
	fmt.Scanln(&mark)
	if mark >= 5 {
		fmt.Println("Excellent!")
	}
}
