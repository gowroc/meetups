package main

import (
	"runtime"
	"time"
)

func runForever() {
	for {
	}
}

func main() {
	for i := 0; i < runtime.GOMAXPROCS(-1); i++ {
		go runForever()
	}
	time.Sleep(time.Second)
}
