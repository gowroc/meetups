package main

import (
	"runtime"
	"time"
)

func runForever() {
	for {
		_ = make([]int, 100000) // HL
	}
}

func main() {
	for i := 0; i < runtime.GOMAXPROCS(-1); i++ {
		go runForever()
	}
	time.Sleep(time.Second)
}
