package main

import (
	"fmt"
	"sync"
)

// Wait group is like a coordinator that helps to Wait for other code to finish
func printNums(wg *sync.WaitGroup, num int) {
	defer wg.Done()
	fmt.Println(num)
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	go printNums(&wg, 1)
	go printNums(&wg, 2)

	wg.Wait()

	fmt.Println("Everything done...")

	/*
		OUTPUT:
			1
			2
			Everything done...
	*/
}
