package main

import (
	"runtime"
	"time"
)

func runForever() {
	for {
		runtime.Gosched()
	}
}

func main() {
	for i := 0; i < runtime.GOMAXPROCS(-1); i++ {
		go runForever()
	}
	time.Sleep(time.Second)
}
