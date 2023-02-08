// package main

// import "fmt"

// func main() {
// fmt.Println("Hello World")
// }

// Declaring variables
package main

import "fmt"

// func main() {
// 	// a := 4
// 	var arr = [...]int{2, 3, 5, 6, 7, 8, 5, 4, 6, 8, 5, 78, 5, 8, 6}
// 	fmt.Print(arr[0:4])
// 	// loops
// 	for i := 0; i < len(arr); i++ {
// 		fmt.Println(arr[i])
// 	}
// }

// structure
type Person struct {
	id     int
	name   string
	age    int
	salary int
}

func main() {
	var p1 Person
	p1.id = 19294
	p1.name = "Subham"
	p1.age = 20
	p1.salary = 294999999

	fmt.Print(p1)

}
