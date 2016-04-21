package main

import "fmt"
import "time"

func function() func() {
	gone := "variable from a no longer available stack frame?" // HL
	return func() {
		fmt.Println(gone) // HL
	}
}

func main() {
	fun := function()
	fun()    // HL
	go fun() // HL
	time.Sleep(time.Millisecond)
}
