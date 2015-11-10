package main

import (
	"fmt"
	"time"
)

func main() {
	done := time.NewTimer(time.Second * 10)
	for {
		select {
		case <-time.After(time.Second * 1):
			fmt.Println("Hello Wroclaw")
		case <-done.C:
			fmt.Println("Quitting")
			return
		}
	}
}
