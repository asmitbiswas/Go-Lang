package main

import "fmt"

func main() {
	// age:= 15

	// if age >= 18 {
	// 	fmt.Println("Person is an adult")
	// } else{
	// 	fmt.Println("Person isn't an adult")
	// }

	age := 112

	if age >= 18 {
		fmt.Println("Person is a adult")
	} else if age >= 12 {
		fmt.Println("Person is not an adult")
	} else {
		fmt.Println("Peson is a kid")
	}
}
