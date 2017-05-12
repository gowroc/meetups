package main

import (
	"fmt"
	"golang.org/x/sys/unix"
	"runtime"
)

// START OMIT
func main() {
	go func() {
		for i := 0; ; i++ {
			fmt.Println(i)
		}
	}()

	for i := 0; i < runtime.GOMAXPROCS(-1); i++ {
		go unix.Pause()
	}

	select {}
}

// END OMIT
