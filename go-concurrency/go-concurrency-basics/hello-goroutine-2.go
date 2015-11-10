package main

import (
	"fmt"
	"sync"
)

func main() {
	// create WaitGroup to sync state between goroutines
	wg := sync.WaitGroup{}
	wg.Add(1) // increment WaitGroup
	go func() {
		fmt.Println("Hello Wroclaw")
		wg.Done() // decrement WaitGroup when job is done
	}()

	wg.Wait() // wait until WaitGroup reaches value 0
}
