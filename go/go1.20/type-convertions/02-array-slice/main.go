package main

import "fmt"

func main() {
	var s = []int{1, 2, 3, 4, 5}
	var a = [6]int(s) // panic: runtime error: cannot convert slice with length 5 to array or pointer to array with length 6
	fmt.Printf("%v -> %v", s, a)
}
