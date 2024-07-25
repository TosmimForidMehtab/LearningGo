package main

import "fmt"

// Variadic: Concept of passing n number of arguments

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	total := sum(1, 2, 3, 4, 5)
	fmt.Println(total)
	intArr := []int{6, 7, 8, 9, 10}
	total = sum(intArr...)
	fmt.Println(total)
}
