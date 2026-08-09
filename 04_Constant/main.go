package main

import "fmt"

// const age = 45 // works here top
// age := 141 /not works here \

func main() {
	// const passKey string = "53***********"

	// const age = 45

	// fmt.Println(age)

	const (
		host = "localhost:"
		port = "3000"
	)
	fmt.Println(host, port)
}
