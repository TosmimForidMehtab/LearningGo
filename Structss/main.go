package main

import "fmt"

// A struct (structure) is a collection of different data types, into a single variable
type Engine struct {
	id     int16
	milage float32
}
type Car struct {
	brand      string
	reg_num    int
	is_working bool
	price      float64
	engine     Engine
}

func printCar(car Car) {
	fmt.Println("Brand ", car.brand)
	fmt.Println("Reg. ", car.reg_num)
	fmt.Println("Status ", car.is_working)
	fmt.Println("Price ", car.price)
	fmt.Println("Engine_id ", car.engine.id)
	fmt.Println("Milage ", car.engine.milage)
	fmt.Println("=========================================")
}

func main() {
	var car1 Car
	car1.brand = "Yotato"
	car1.reg_num = 12345
	car1.is_working = true
	car1.price = 220000
	car1.engine.id = 11
	car1.engine.milage = 22
	printCar(car1)

	car2 := Car{
		brand:      "ATAT",
		reg_num:    67890,
		is_working: true,
		price:      300000,
		engine: Engine{
			id:     21,
			milage: 30,
		},
	}
	printCar(car2)

	// Anonymous structs
	u := struct {
		firstName string
		lastName  string
	}{
		firstName: "John",
		lastName:  "Doe",
	}
	fmt.Println(u.firstName, " ", u.lastName)
}
