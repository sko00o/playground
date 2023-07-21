package main

import (
	"fmt"
	"sync/atomic"
	"unsafe"
)

type T struct {
	_ [0]atomic.Int64
	x uint64
}

func main() {
	t := T{}
	fmt.Printf("size = %d, addr = %p, aligned = %v\n",
		unsafe.Sizeof(t),
		&t,
		uintptr(unsafe.Pointer(&t))%64 == 0)
}
