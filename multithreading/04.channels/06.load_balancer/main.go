package main

import (
	"fmt"
	"time"
)

func worker(workerId int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d received: %d\n", workerId, x)
		time.Sleep(time.Second)
	}
}

func main() {
	data := make(chan int)
	qtyWorkers := 50_000

	// Initialize workers
	for i := 0; i < qtyWorkers; i++ {
		go worker(i, data)
	}

	for i := 0; i < 10_000_000; i++ {
		data <- i
	}
}
