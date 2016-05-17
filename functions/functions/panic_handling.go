package main

import "fmt"

func handle() {
	msg := recover() // HL
	fmt.Println("handled", msg)
}

func main() {
	defer handle() // HL
	panic("error")
}
