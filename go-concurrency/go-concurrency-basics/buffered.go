package main

import "fmt"

func main() {
	out := make(chan bool, 1) // make buffered channel with capacity 1

	out <- true // send true value to channel

	fmt.Println("I am done!")
}
