package main

import "fmt"

func main() {
	var myVar interface{} = "Waldrey Souza Silva"

	fmt.Println(myVar.(string))
	res, ok := myVar.(int)
	fmt.Printf("O resultado de res é %v e o resultado de ok é %v\n", res, ok)

	res2 := myVar.(int)
	fmt.Printf("O result de res2 é %v\n", res2)
}
