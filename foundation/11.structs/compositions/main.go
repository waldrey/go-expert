package main

import "fmt"

type Address struct {
	Street string
	Number int
	City   string
	State  string
}

type Client struct {
	Name   string
	Age    int
	Active bool
	Address
	// Endereco Address
}

func main() {
	waldrey := Client{
		Name:   "Waldrey",
		Age:    25,
		Active: true,
	}

	waldrey.Active = false
	waldrey.State = "São Paulo"
	waldrey.Address.City = "São José dos Campos"

	fmt.Printf("Name: %s | Age: %d | Is active?: %t\n", waldrey.Name, waldrey.Age, waldrey.Active)
}
