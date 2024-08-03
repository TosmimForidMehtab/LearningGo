package main

import (
	"fmt"
	"strings"
)

func main() {
	// Replace a string with another
	s := "Hello World"
	replacer := strings.NewReplacer("World", "Tos")
	s = replacer.Replace(s)
	fmt.Println(s)      // Hello Tos
	fmt.Println(len(s)) // 9

	// Check if a string contains another string
	if strings.Contains(s, "Hello") {
		fmt.Println("Hello is in the string") // This is executed and printed
	} else {
		fmt.Println("Hello is not in the string")
	}

	// Find the first occuring index of a string
	fmt.Println(strings.Index(s, "o")) // 4

	// Search and replace every o with 0
	fmt.Println(strings.Replace(s, "o", "0", -1)) // Hell0 T0s

	// Split string
	fmt.Println(strings.Split(s, " ")) // [Hello Tos]

	// Uppercase and lowercase
	fmt.Println(strings.ToUpper(s)) // HELLO TOS
	fmt.Println(strings.ToLower(s)) // hello tos

	// Check if a string starts with another string
	fmt.Println(strings.HasPrefix(s, "He")) // true

	// Check if a string ends with another string
	fmt.Println(strings.HasSuffix(s, "ld")) // false
}
