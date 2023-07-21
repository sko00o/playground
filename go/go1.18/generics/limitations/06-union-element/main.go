package main

type MyInterface interface {
	M()
}

type GenericsInterface interface {
	float64 | ~int | MyInterface // cannot use main.MyInterface in union (main.MyInterface contains methods)
}

func main() {}
