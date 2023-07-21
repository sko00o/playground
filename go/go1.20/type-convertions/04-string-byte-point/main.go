package main

import (
	"fmt"
	"unsafe"
)

func main() {
	s := "hello!"
	b := unsafe.StringData(s)
	*b = 'j' // fatal error: fault unexpected fault address 0x499982
	fmt.Println(s)
}
