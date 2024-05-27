package main

func sum(a, b *int) int {
	*a = 50
	*b = 50
	return *a + *b
}

func main() {
	myVar1 := 10
	myVar2 := 20
	println(myVar1)
	println(myVar2)
	println(sum(&myVar1, &myVar2))
	println(myVar1)
	println(myVar2)
}
