package main

import "time"

func main() {

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		println("It's the weekend")

	default:
		println("It's a weekday")
	}

	// Type switch
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			println("I'm a bool")
		case int:
			println("I'm an int")
		case string:
			println("I'm a string")
		default:
			println("Other", t)
		}
	}
	whatAmI(true)
	whatAmI("hello")
	whatAmI(23.5)

}
