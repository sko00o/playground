package main

func main() {
	GenericFoo(T0{})
}

func GenericFoo[P PType](x P) {
	x.m()
	x.M() // x.M undefined (type P has no field or method M, but does have m)
}

type PType interface {
	T0 | T1
	m()
}

type T0 struct{}

func (T0) m() {}
func (T0) M() {}

type T1 struct{}

func (T1) m() {}
func (T1) M() {}
