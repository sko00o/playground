package main

func main() {
	GenericFoo(T0{})
}

func GenericFoo[P PType](x P) {
	_ = x.F // x.F undefined (type P has no field or method F)
	_ = x.f // x.f undefined (type P has no field or method f)
}

type PType interface {
	T0 | T1
}

type T0 struct {
	f int
	F int
}

type T1 struct {
	f int
	F int
}
