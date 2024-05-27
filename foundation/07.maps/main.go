package main

import "fmt"

func main() {
	salarios := map[string]int{"Waldrey": 16000, "Bruno": 4600, "Alan": 2700}

	// sal := make(map[string]int)
	// sal1 := map[string]int{}
	// sal1["Waldrey"] = 1000

	for nome, salario := range salarios {
		fmt.Printf("O salário de %s é de %d\n", nome, salario)
	}

	for _, salario := range salarios {
		fmt.Printf("O salário é de %d\n", salario)
	}
}
