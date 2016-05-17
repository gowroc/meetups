package main

import (
	"fmt"
	"sync"
)

func do(wg *sync.WaitGroup, i int) {
	defer wg.Done() // HL
	fmt.Println(i)
}

func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go do(wg, i)
	}
	wg.Wait()
}
