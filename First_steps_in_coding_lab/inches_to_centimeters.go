// package main

// import "fmt"

// func main() {
// 	var inch float32 = 0.39
// 	var cm float32
// 	fmt.Scanln(&cm)
// 	var result = cm * inch
// 	fmt.Println(result)
// }

package main

import "fmt"

func main() {
	var inches float32
	fmt.Scanln(&inches)
	var cm float32 = inches * 2.54
	fmt.Println(cm)

}
