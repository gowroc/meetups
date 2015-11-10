package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Hello Wroclaw")
	fmt.Printf("Current number of goroutines: %d \n", runtime.NumGoroutine())
}
