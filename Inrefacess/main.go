package main

import "fmt"

// Interfaces define a contract for what methods a type must have. They don't provide the implementation for those methods rather the implementation is provided by concrete types that satisfy the Interface
type Speaker interface {
	Speak() string
}

type Cat struct{}

func (c Cat) Speak() string {
	return "Meoww!"
}

func speak(speaker Speaker) {
	fmt.Println("Speaker says ", speaker.Speak())
}

func main() {
	speak(Cat{}) // Speaker says  Meoww!
}
