package main

import "runtime"

func main() {
	runtime.GOMAXPROCS(1)
	c := make(chan bool)
	go func() {
		<-c
		for {
			runtime.Gosched() // HL
		}
	}()
	c <- true
}
