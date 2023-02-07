// package main

// import "fmt"

// func main() {
// fmt.Println("Hello World")
// }

// Declaring variables
package main

import "fmt"

func main() {
	// a := 4
	var arr = [...]int{2, 3, 5, 6, 7, 8, 5, 4, 6, 8, 5, 78, 5, 8, 6}
	fmt.Print(arr[0:4])
	// loops
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}

//
