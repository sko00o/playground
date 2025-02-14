package main

type P[K comparable, V any] struct {
	next *K
	K    // embedded field type cannot be a (pointer to a) type parameter
	*V   // embedded field type cannot be a (pointer to a) type parameter
}

type iterable interface{}

type Scanner[item iterable] interface {
	Next() item
	item // cannot embed a type parameter
}

func main() {
	var s Scanner[int]
	_ = s
}
