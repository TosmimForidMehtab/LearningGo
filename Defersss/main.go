package main

import "fmt"

// Defer keyword is used to delay the execution of a function or a statement until the newrby function returns.

func greet() {
	defer fmt.Println("Hello, World!") // This will be executed at last
	fmt.Println("This is a greeting")
}

func main() {
	greet()
}
