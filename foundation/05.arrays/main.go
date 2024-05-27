package main

import "fmt"

const a = "Hello World!"

type ID int

var (
	b bool    = true
	c int     = 10
	d string  = "Waldrey"
	e float64 = 1.2
	f ID      = 5
)

func main() {
	var myArray [3]int
	myArray[0] = 10
	myArray[1] = 20
	myArray[2] = 30

	for i, v := range myArray {
		fmt.Printf("O valor do índice é %d e o valor é %d", i, v)
	}
}
