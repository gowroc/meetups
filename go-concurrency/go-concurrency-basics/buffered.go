package main

import "fmt"

func main() {
	out := make(chan bool, 1)

	out <- true

	fmt.Println("I am done!")
}
