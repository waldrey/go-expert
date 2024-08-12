package main

import "fmt"

// Thread 1
func main() {
	ch := make(chan int)

	go publish(ch)
	reader(ch)

}

// Thread 2
func publish(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

// Thread 3
func reader(ch chan int) {
	for x := range ch {
		fmt.Printf("Received: %d\n", x)
	}
}
