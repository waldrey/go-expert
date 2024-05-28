package main

import (
	"fmt"
	"go-expert/matematica"
	"github.com/google/uuid"
)

func main() {
	s := matematica.Soma(10, 20)

	carro := matematica.Carro{Marca: "Toyota"}
	fmt.Println(carro.Acelerar())


	fmt.Printf("O resultado: %v\n", s)
	fmt.Println(matematica.A)
	fmt.Println(uuid.New())
}
