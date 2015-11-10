package main

import (
	"fmt"
)

func main() {
	done := make(chan bool) // create unbuffered channel of bools
	go func() {
		fmt.Println("Hello Wroclaw")
		done <- true // send value when job is done
	}()

	<-done // wait for first message
}
