package main

import (
	"fmt"
	"runtime"
)

// START OMIT
func Pause()

func main() {
	go func() {
		for i := 0; ; i++ {
			fmt.Println(i)
		}
	}()

	for i := 0; i < runtime.GOMAXPROCS(-1); i++ {
		go Pause()
	}

	select {}
}

// END OMIT
