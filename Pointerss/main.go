package main

import "fmt"

func main() {
	var ptr *int
	fmt.Println("Val: ", ptr) // <nil>

	myNumber := 10
	var ptr2 = &myNumber
	var ptr3 *int = &myNumber
	fmt.Println("Val of ptr2:", *ptr2)
	fmt.Println("Val of ptr3:", *ptr3)
}
