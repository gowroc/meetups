package main

import "flag"
import "fmt"
import "github.com/gowroc/meetups/kafka2go/kafka"

var groupName string

func main() {
	flag.StringVar(&groupName, "group", "group", "consumer group name")
	flag.Parse()

	consumer, err := kafka.NewConsumer(groupName, "example", "localhost:2181")
	if err != nil {
		panic(fmt.Sprintf("Failed to initialize Kafka consumer: %v", err))
	}

	err = consumer.Start()
	if err != nil {
		panic(fmt.Sprintf("Failed to start Kafka consumer: %v", err))
	}
}
