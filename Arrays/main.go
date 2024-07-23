package main

import "fmt"

func main() {
	var a = [5]int{1, 2, 3, 4}

	fmt.Println(a)      // [1 2 3 4 0] The last item is 0 as it is not initialized
	fmt.Println(len(a)) // 5

	var b [5]string
	b[0] = "Hello"
	b[1] = "World"
	b[2] = "!"
	b[3] = "!"

	fmt.Println(b)        // [Hello World ! ! ] The last item is an empty space as it is not initialized
	fmt.Println(len(b))   // 5
	fmt.Printf("%T\n", b) // [5]string

	// 2D array
	nums := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(nums) // [[1 2] [3 4]]

}
