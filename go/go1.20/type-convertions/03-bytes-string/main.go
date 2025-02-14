package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var bs = []byte{'h', 'e', 'l', 'l', 'o', '!'}
	// var bs = [6]byte{'h', 'e', 'l', 'l', 'o', '!'}
	s := unsafe.String(&bs[0], len(bs))
	fmt.Println(s) // hello!
	bs[0] = 'j'
	fmt.Println(s) // jello!
}
