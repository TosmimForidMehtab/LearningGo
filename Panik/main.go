package main

import "fmt"

// Panic is used to raise exception by the programmer or due to some unexpected errors
func isEligible(age int) {
	if age > 18 {
		fmt.Println("You are eligible!")
	} else {
		panic("Paanikk! Noooooooooooooooooooooo!")
	}
}
func main() {
	isEligible(20)
	isEligible(10)
}
