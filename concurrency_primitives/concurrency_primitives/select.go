package main

import (
	"fmt"
	"time"
)

func generate(n int) <-chan int {
	out := make(chan int)
	go func() {
		for {
			out <- n
			time.Sleep(time.Second / 10)
		}
	}()
	return out
}

func main() {
	zero, one := generate(0), generate(1)

	for {
		select {
		case a := <-zero: // HL
			fmt.Print(a)
		case b := <-one: // HL
			fmt.Print(b)
		}
	}
}
