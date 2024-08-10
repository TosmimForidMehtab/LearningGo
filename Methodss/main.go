package main

import "fmt"

// A method is a function associated with a particular type. It is a way to define behavior for specific type.
type Person struct {
	first_name string
	last_name  string
}

func (p Person) getFullName() string {
	return p.first_name + " " + p.last_name
}
func main() {
	person := Person{first_name: "John", last_name: "Doe"}
	fmt.Println(person.getFullName())
}
