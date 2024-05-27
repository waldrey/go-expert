package main

func main() {
	a := 10

	a = 20
	println(a)
	var pointer *int = &a
	*pointer = 20

	b := &a
	*b = 30
	println(a)
}
