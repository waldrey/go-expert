package main

import "fmt"

// Thread 1
func main() {
	channel := make(chan string) // Empty channel

	// Thread 2
	go func() {
		channel <- "Hello World!" // Full channel
	}()

	// Thread 1
	msg := <-channel // Empty channel
	fmt.Println(msg)
}
