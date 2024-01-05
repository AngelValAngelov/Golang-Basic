// package main

// import "fmt"

//	func main() {
//		var length, width int
//		fmt.Scanln(&length)
//		fmt.Scanln(&width)
//		print(length * width)
//	}
package main

import "fmt"

func main() {
	var a, b int
	fmt.Scanln(&a)
	fmt.Scanln(&b)

	var area int = a * b
	fmt.Println(area)

}
