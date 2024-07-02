package main

import "github.com/waldrey/golang/packaging/mod_replace/math"

func main() {
	m := math.NewMath(1, 2)
	println(m.Add())
}
