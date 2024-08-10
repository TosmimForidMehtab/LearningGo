package main

import (
	"fmt"
	"time"
)

// The select statement is used to handle multiple channels in a non-blocking way. It provides a way to wait on multiple communication operations simultaneously and execute the first case that is ready
func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	// We use the select statement to wait on both channels. If both channels are ready, the
	// select statement will execute the first case that is ready. If only one channel is ready,
	// the select statement will execute the case for that channel. If neither channel is ready,
	// the select statement will block until one of the channels is ready.

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Channel1 message"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "Channel2 message"
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println("Received message from channel 1:", msg1)
	case msg2 := <-ch2:
		fmt.Println("Received message from channel 2:", msg2)
	case <-time.After(5 * time.Second):
		fmt.Println("Timeout occurred")
	}
}
