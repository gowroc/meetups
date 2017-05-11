package main

import (
	"fmt"
	"runtime"
)

// START OMIT

// #include <unistd.h>
import "C"

func main() {
	go func() {
		for i := 0; ; i++ {
			fmt.Println(i)
		}
	}()

	for i := 0; i < runtime.GOMAXPROCS(-1); i++ {
		go C.pause()
	}

	select {}
}

// END OMIT
