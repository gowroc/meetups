package main

import "fmt"
import "time"

func receive(messages chan string) {
	for {
		fmt.Println("received", <-messages) // HL
	}
}

func send(messages chan string) {
	for i := 0; i < 10; i++ {
		messages <- fmt.Sprintf("%d", i) // HL
		fmt.Println("send", i)
	}
}

func main() {
	messages := make(chan string, 3) // HL
	go receive(messages)
	go send(messages)
	time.Sleep(time.Millisecond)
}
