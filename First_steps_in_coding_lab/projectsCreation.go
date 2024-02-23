package main

import "fmt"

func main() {
	var architect string
	fmt.Scanln(&architect)

	var projectsCount int
	fmt.Scanln(&projectsCount)

	necessaryHours := projectsCount * 3

	fmt.Printf("The architect %s will need %d hours to complete %d project/s.", architect, necessaryHours, projectsCount)

}
