package main

import (
	"fmt"
	"time"
)

func main() {
	for val := range []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10} {
		go func() {
			fmt.Println(val)
		}()
	}
	time.Sleep(time.Second)
}
