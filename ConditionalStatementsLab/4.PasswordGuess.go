package main

import "fmt"

func main() {
	password := "s3cr3t!P@ssw0rd"

	var passwordInput string
	fmt.Scanln(&passwordInput)

	if passwordInput == password {
		fmt.Println("Welcome")
	} else {
		fmt.Println("Wrong password!")
	}
}
