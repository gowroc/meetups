package main

import "fmt"

func main() {
	done := make(chan bool)
	go func() {
		fmt.Println("x")
	}()
	<-done // HL
}
