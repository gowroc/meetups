package main

import "fmt"
import "time"

func async(messages chan string) {
	text := <-messages // HL
	fmt.Println(text)
}

func main() {
	messages := make(chan string)
	go async(messages)
	messages <- "hello" // HL
	time.Sleep(time.Millisecond)
}
