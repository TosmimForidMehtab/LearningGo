package main

import "fmt"

func main() {
	const name string = "John"
	// name = "Joe" // Not allowed
	fmt.Println(name)

	// Grouping multiple constants
	const (
		pi       = 3.14
		language = "Go"
	)
	fmt.Println(pi, language)
}
