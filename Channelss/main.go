package main

import "fmt"

// Channels are like communication pipes between different parts of your program. They let different pieces of code talk to each other by sending and receiving message
func sendInt(ch chan<- int) {
	ch <- 1
}

func main() {
	// Create a channel
	ch := make(chan int)
	go sendInt(ch) // If `go` not used, it will give error
	val := <-ch
	fmt.Println(val)

	// Unbuffered channel
	unbuff := make(chan int)
	// Buffered channel
	buff := make(chan string, 2)

	go sendInt(unbuff)
	unbuffVal := <-unbuff
	fmt.Println("Value from Unbuffered channel", unbuffVal)

	go func() {
		buff <- "Hello"
		buff <- "World"
		close(buff)
	}()
	// If we don't close the channel, it will block indefinitely
	for msg := range buff {
		fmt.Println("Data from Buffered channel", msg)
	}
}
