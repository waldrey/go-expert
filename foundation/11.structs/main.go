package main

import "fmt"

type Client struct {
	Name   string
	Age    int
	Active bool
}

func main() {
	waldrey := Client{
		Name:   "Waldrey",
		Age:    25,
		Active: true,
	}

	waldrey.Active = false

	fmt.Printf("Name: %s | Age: %d | Is active?: %t\n", waldrey.Name, waldrey.Age, waldrey.Active)
}
