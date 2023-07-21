package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	t1()
	t2()
}

func t1() {
	var n uint64
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Printf("g%02d: n+1 = %d\n", i, atomic.AddUint64(&n, 1))
		}(i)
	}
	wg.Wait()
	fmt.Printf("finally n = %d\n", atomic.LoadUint64(&n))
}

func t2() {
	var n atomic.Uint64
	// n.Store(0) // option: set initial value
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Printf("g%02d: n+1 = %d\n", i, n.Add(1))
		}(i)
	}
	wg.Wait()
	fmt.Printf("finally n = %d\n", n.Load())
}
