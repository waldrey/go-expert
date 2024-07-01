package main

import "fmt"

func panic1() {
	panic("panic1")
}

func panic2() {
	panic("panic2")
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			if r == "panic1" {
				fmt.Println("Panic 1 recovered:", r)
			}

			if r == "panic2" {
				fmt.Println("Panic 2 recovered:", r)
			}

		}
	}()

	panic2()
}
