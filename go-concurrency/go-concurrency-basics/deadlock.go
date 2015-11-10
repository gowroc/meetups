package main

import "fmt"

func main() {
	out := make(chan bool) // make unbuffered channel

	// send true value into the channel
	// there is no one who could reads from this channel
	// this causes deadlock
	out <- true

	fmt.Println("I am done!")
}
