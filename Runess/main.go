package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "abcdefg"
	fmt.Println(utf8.RuneCountInString(str)) // 7

	// Printing the index, utf8 code and character
	for idx, val := range str {
		fmt.Printf("%d : %#U : %c\n", idx, val, val)
	}
	// OUTPUT:
	/*
		0 : U+0061 'a' : a
		1 : U+0062 'b' : b
		2 : U+0063 'c' : c
		3 : U+0064 'd' : d
		4 : U+0065 'e' : e
		5 : U+0066 'f' : f
		6 : U+0067 'g' : g
	*/
}
