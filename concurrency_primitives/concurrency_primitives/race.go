package main

import (
	"fmt"
	"sync"
)

func inc(x *int, wg *sync.WaitGroup) {
	*x++ // HL
	wg.Done()
}

func main() {
	x := 0
	wg := &sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go inc(&x, wg) // HL
	}
	wg.Wait()
	fmt.Println(x)
}
