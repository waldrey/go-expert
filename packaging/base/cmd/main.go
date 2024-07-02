package main

import (
	"fmt"

	"github.com/waldrey/golang/packaging/base/math"
)

func main() {
	m := math.NewMath(1, 2)
	m.C = 5
	fmt.Println(m.C)
	// fmt.Println(m.Add())
	// fmt.Println("Hello world!")
}
