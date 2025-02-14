package main

func main() {
	var i = complex(1.0, 2.0) // 1+2i
	GenericsFoo(i)
}

type ComplexType interface {
	~complex64 | ~complex128
}

func GenericsFoo[T ComplexType](c T) T {
	r := real(c) // c (variable of type T constrained by Complex) not supported as argument to real for go1.18 (see issue #50937)
	println(r)

	i := imag(c) // c (variable of type T constrained by Complex) not supported as argument to real for go1.18 (see issue #50937)
	println(i)

	return complex(r, i) // invalid operation: complex(r, i) (mismatched types T and T)
}
