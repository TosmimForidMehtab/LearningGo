package main

import "fmt"

func main() {
	// For loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// There is no while loop construct in go
	// Instead we can use for loop
	i := 5
	for i < 10 {
		fmt.Println(i)
		i++
	}

	// For loop with break and continue. Once we reach 5, the looping will stop and we will skip 3
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		if i == 3 {
			continue
		}
		fmt.Println("Num -> ", i)
	}

	// Range bases loop
	arr := []int{1, 2, 3, 4, 5}
	for index, value := range arr {
		fmt.Println(index, value)
	}

	for ele := range 5 {
		fmt.Println(ele) // 0 1 2 3 4
	}

	// Infinite loop
	// for {
	// 	fmt.Println("Infinite loop")
	// }

}
