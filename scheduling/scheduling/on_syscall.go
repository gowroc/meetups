package main

import (
	"fmt"
	"runtime"
	"time"
)

func runForever() {
	for {
		fmt.Print("") // HL
	}
}

func main() {
	for i := 0; i < runtime.GOMAXPROCS(-1); i++ {
		go runForever()
	}
	time.Sleep(time.Second)
}
