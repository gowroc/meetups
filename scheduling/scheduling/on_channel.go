package main

import (
	"runtime"
	"time"
)

func runForever() {
	bc := make(chan bool) // HL
	for {
		bc <- true // HL
	}
}

func main() {
	for i := 0; i < runtime.GOMAXPROCS(-1); i++ {
		go runForever()
	}
	time.Sleep(time.Second)
}
