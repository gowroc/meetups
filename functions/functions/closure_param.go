package main

import (
	"fmt"
	"time"
)

func main() {
	data := []int{1, 2, 3}
	for i, v := range data {
		go func() {
			fmt.Println(i, v)
		}()
	}
	time.Sleep(time.Millisecond)
}
