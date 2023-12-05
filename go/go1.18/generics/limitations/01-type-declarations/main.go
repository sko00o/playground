package main

func main() {
	GenericFoo(10)
}

func GenericFoo[T any](t T) {
	type MyInt int // type declarations inside generic functions are not currently supported
	var i MyInt = 10
	println(i)
}
