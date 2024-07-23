package main

import (
	"fmt"
	"slices"
	"sort"
)

func main() {
	// Not providing the size makes it a slice
	var a = []string{}
	fmt.Printf("%T", a) // []string
	fmt.Println()
	a = append(a, "Hello")
	a = append(a, "World")
	a = append(a, "!")
	fmt.Println(a) // [Hello World !]

	b := append(a[1:2]) // slicing similar to python
	fmt.Println(b)      // [World]

	// Way 2 of creating a slice
	c := make([]int, 3) // c initialized with 3 elements with all values 0
	c[0] = 1
	c[1] = 3
	c[2] = 2
	// c[3] = 4 // out of range error
	c = append(c, 4, 5, 6) // No error
	fmt.Println(c)

	if !sort.IntsAreSorted(c) {
		sort.Ints(c)
		fmt.Println("Sorting...")
	}
	fmt.Println(c) // Sorted slice

	// Removing an element
	indexToRemove := 2
	c = append(c[:indexToRemove], c[indexToRemove+1:]...)
	fmt.Println(c) // [1 2 4 5 6]

	// Capacity
	fmt.Println(cap(c)) // 6

	// Copying a slice
	arr := []int{}
	arr = append(arr, 1, 2, 3, 4, 5)
	arr2 := make([]int, len(arr))

	copy(arr2, arr)
	fmt.Println(arr, arr2) // [1 2 3 4 5] [1 2 3 4 5]

	// Check if 2 slices are same
	// fmt.Println(arr == arr2) // This won't work
	fmt.Println(slices.Equal(arr, arr2)) // true

	// 2D slice
	arr3 := [][]int{{1, 2}, {3, 4}}
	fmt.Println(arr3) // [[1 2] [3 4]]

}
