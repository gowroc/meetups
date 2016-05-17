package main

import "fmt"

func main() {
	defer fmt.Println("hello") // HL
	panic("error")
}
