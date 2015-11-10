package main

import (
	"fmt"
	"runtime"
)

func main() {
	done := make(chan bool)
	go func() {
		fmt.Println("Hello Wroclaw")
		done <- true
	}()

	fmt.Printf("Current number of goroutines: %d \n", runtime.NumGoroutine())
	<-done

	fmt.Printf("Current number of goroutines: %d \n", runtime.NumGoroutine())
}
