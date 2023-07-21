package main

func main() {
	t1()
	t2()
}

func t1() {
	var a interface{} = int(1)
	var b interface{} = int(2)
	println(isEqual(a, b)) // Go 1.18: interface{} does not implement comparable
}

func t2() {
	var a interface{} = []byte{1}
	var b interface{} = []byte{2}
	println(isEqual(a, b)) // Go 1.20: panic: runtime error: comparing uncomparable type []uint8
}

func isEqual[T comparable](a, b T) bool {
	return a == b
}
