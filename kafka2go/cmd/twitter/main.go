package main

import (
	"fmt"
	"time"

	"github.com/gowroc/meetups/kafka2go/kafka"
	"github.com/gowroc/meetups/kafka2go/twitter"
)

// Tweet is a representation of Twitter's message
type Tweet struct {
	Time time.Time
	Text string
}

func main() {
	var (
		lastTweetID uint64
		handle      = "mycodesmells"
	)

	client := twitter.Connect()
	producer, err := kafka.NewProducer("0.0.0.0:9092")
	if err != nil {
		panic(fmt.Sprintf("Failed to initialize Kafka producer: %v", err))
	}

	for {
		select {
		case <-time.After(time.Second):
			tl, err := twitter.Tweets(client, handle, lastTweetID)
			if err != nil {
				fmt.Printf("Failed to load Tweets from %s: %v", handle, err)
			}

			for _, t := range tl {
				tweet := Tweet{Time: t.CreatedAt(), Text: t.Text()}
				producer.SendMessage("example", tweet)
			}

			fmt.Printf("Found %d tweets\n", len(tl))

			if len(tl) > 0 {
				lastTweetID = tl[0].Id()
				tl[0].CreatedAt()
			}
		}
	}

}
