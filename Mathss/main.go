package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	// Random number
	seedSecs := time.Now().Unix()
	rand.NewSource(seedSecs)
	randNum := rand.Intn(10) + 1
	fmt.Println("Random :", randNum)

	// Absolute of a number
	fmt.Println("Absolute of -5 is:", math.Abs(-5)) // 5

	// Square root of a number
	fmt.Println("Square root of 25 is:", math.Sqrt(25)) // 5

	// Power of a number
	fmt.Println("2 to the power of 3 is:", math.Pow(2, 3)) // 8

	// Floor of a number
	fmt.Println("Floor of 3.7 is:", math.Floor(3.7)) // 4

	// Ceil of a number
	fmt.Println("Ceil of 3.7 is:", math.Ceil(3.7)) // 3

	// Min and max of numbers
	fmt.Println("Min of 10 and 20 is:", math.Min(10, 20)) // 10
	fmt.Println("Max of 10 and 20 is:", math.Max(10, 20)) // 20

	// Logarithm of a number base 10
	fmt.Println("Logarithm of 10 is:", math.Log10(10)) // 1

	// Logarithm of a number base 2
	fmt.Println("Logarithm of 8 is:", math.Log2(8)) // 3
}
