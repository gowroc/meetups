package main

import (
	"fmt"
	"sync"
)

func inc(m map[int]int, i int, wg *sync.WaitGroup) {
	m[i] = i * i // HL
	wg.Done()
}

func main() {
	m := make(map[int]int)
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go inc(m, i, wg) // HL
	}
	wg.Wait()
	fmt.Println(m)
}
