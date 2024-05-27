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
	fmt.Printf("O tipo de E é %T\n", e)
	fmt.Printf("O valor de E é %v", e)
}
