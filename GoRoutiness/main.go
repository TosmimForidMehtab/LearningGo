package main

import (
	"fmt"
	"sync"
	"time"
)

// Goroutine is a lightweight, independently running thread that allows concurrent execution in a program.
func printNums(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

func main() {
	fmt.Println("Welcome")
	var wg sync.WaitGroup
	wg.Add(1)
	go printNums(&wg)
	wg.Wait()
}
