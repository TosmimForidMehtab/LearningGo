package main

import "fmt"

// a := 1 // short variable declaration not allowed in the global scope
// Capital letter is used for public variables

func main() {
	var username string = "Tosmim"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isTrue bool = true
	fmt.Println(isTrue)

	var smallF float32 = 255.37747474
	var largeF float64 = 255.37747474747456
	fmt.Println(smallF)
	fmt.Println(largeF)

	// Default values and aliases
	var x int
	fmt.Println(x) // 0 not garbage value
	fmt.Printf("Variable is of type %T", x)
	fmt.Println()

	// Implicit type
	var str = "Hello"
	fmt.Println(str)

	// No var type
	y := 3000 // := is called the walrus/short assignment operator.
	fmt.Println(y)
}
