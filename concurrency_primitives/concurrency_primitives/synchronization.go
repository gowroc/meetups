package main

import "fmt"

func print(do, done chan struct{}) {
	<-do // HL
	fmt.Println("synchronized")
	done <- struct{}{} // HL
}

func main() {
	do := make(chan struct{})
	done := make(chan struct{})
	go print(do, done)

	fmt.Println("before")
	do <- struct{}{} // HL
	<-done           // HL
	fmt.Println("after")
}
