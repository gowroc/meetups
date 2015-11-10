package main

import (
	"fmt"
	"time"
)

func main() {
	done := time.NewTimer(time.Second * 10) // create channel which returns value after 10 second

	// start infinite loop
	for {
		select {
		case <-time.After(time.Second * 1): // put value in channel every second
			fmt.Println("Hello Wroclaw")
		case <-done.C: // wait for message from done channel
			fmt.Println("Quitting")
			return
		}
	}
}
