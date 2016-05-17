package main

import (
	"fmt"
	"time"
)

func print(text string) {
	for {
		fmt.Print(text)
		time.Sleep(time.Second / 10)
	}
}

func main() {
	go print("0") // HL
	go print("1") // HL
	time.Sleep(5 * time.Second)
}
