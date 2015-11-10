package main

import "fmt"

func main() {
	out := make(chan bool)

	out <- true

	fmt.Println("I am done!")
}
