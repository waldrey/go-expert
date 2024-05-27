package main

import "fmt"

type Person interface {
	Disable()
}

type Company struct {
	Name string
}

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
}

func (e Company) Disable() {

}

func (c Client) Disable() {
	c.Active = false
	fmt.Printf("O cliente %s foi desativado\n", c.Name)
}

func Deactivation(person Person) {
	person.Disable()
}

func main() {
	waldrey := Client{
		Name:   "Waldrey",
		Age:    25,
		Active: true,
	}

	myCompany := Company{}
	Deactivation(myCompany)

	Deactivation(waldrey)
}
