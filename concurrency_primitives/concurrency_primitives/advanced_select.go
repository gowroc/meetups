package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sender() <-chan struct{} {
	out := make(chan struct{})
	go func() {
		for {
			out <- struct{}{}
			time.Sleep(time.Second / time.Duration(rand.Intn(3)+1))
		}
	}()
	return out
}

func receiver() chan struct{} {
	out := make(chan struct{})
	go func() {
		for {
			<-out
			time.Sleep(time.Second / time.Duration(rand.Intn(3)+1))
		}
	}()
	return out
}

func main() {
	input := sender()
	output := receiver()

	for {
		select {
		case <-input: // HL
			fmt.Print(0)
		case output <- struct{}{}: // HL
			fmt.Print(1)
		default: // HL
			fmt.Print(" ")
		}
		time.Sleep(time.Second / 10)
	}
}
