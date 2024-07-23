package main

import (
	"fmt"
	"maps"
)

func main() {
	// Declaring a map
	m := make(map[string]int)
	// OR
	// m := map[string]int{"k0": 0}

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)       // map: map[k1:7 k2:13]
	fmt.Println("len:", len(m))  // len: 2
	fmt.Println("k1: ", m["k1"]) // 7
	// Deleting map entry
	delete(m, "k2")
	fmt.Println(m)

	// Check if key exists
	val, ok := m["k1"]
	if ok {
		fmt.Println("Key exists", val)
	} else {
		fmt.Println("Key does not exist")
	}

	// Check if 2 maps are equal
	m1 := map[string]int{"foo": 1, "bar": 2}
	m2 := map[string]int{"foo": 1, "bar": 2}
	// fmt.Println(m1 == m2) // This won't work
	fmt.Println(maps.Equal(m1, m2)) // true

	// Iterating over map
	for k, v := range m1 {
		fmt.Println(k, "->", v)
	}

	// If key does not exist in the map, it returns zeroed value
	// CLearing the map
	clear(m)
	fmt.Println(m) // map[]

}
