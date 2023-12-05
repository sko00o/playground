package main

import "fmt"

func main() {
	var s = []int{1, 2, 3, 4, 5}
	var a = [5]int(s) // *(*[5]int)(s)
	var p = (*[5]int)(s)
	fmt.Println("a=", a)
	fmt.Println("p=", p)
	s[2] = 33
	fmt.Println("a=", a)
	fmt.Println("p=", p)

	// Output:
	// a= [1 2 3 4 5]
	// p= &[1 2 3 4 5]
	// a= [1 2 3 4 5]
	// p= &[1 2 33 4 5]
}
