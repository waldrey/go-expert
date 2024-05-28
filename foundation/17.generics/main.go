package main

import "fmt"

type MyNumber int

type Number interface {
	~int | float64
}

func sum[T Number](m map[string]T) T {
	var sum T
	for _, v := range m {
		sum += v
	}

	return sum
}

func sumInteger(m map[string]int) int {
	var sum int
	for _, v := range m {
		sum += v
	}

	return sum
}

func sumFloat(m map[string]float64) float64 {
	var sum float64
	for _, v := range m {
		sum += v
	}

	return sum
}

func compare[T comparable](a T, b T) bool {
	if a == b {
		return true
	}

	return false
}

func main() {
	m := map[string]int{"Waldrey": 13000, "Joao": 4500, "Francesco": 18000}
	m2 := map[string]float64{"Waldrey": 13000.50, "Joao": 4500.22, "Francesco": 18000.25}
	m3 := map[string]MyNumber{"Waldrey": 13000, "Joao": 4500, "Francesco": 18000}

	fmt.Println(sum(m))
	fmt.Println(sum(m2))
	fmt.Println(sum(m3))
	fmt.Println(compare(10, 10.0))
}
