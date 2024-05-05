package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome nigger"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")

	// Comma ok || error ok
	input, err := reader.ReadString('\n')
	fmt.Println("Hello", input)

	var a, b int
	fmt.Println("Enter 2 numbers: ")
	fmt.Scan(&a, &b)
	fmt.Println(a, b)
	fmt.Printf("Type of a, b is %T %T", a, b)

	if err != nil {
		fmt.Println("An error occured", err)
	}
}
