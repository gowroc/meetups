package kafka

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/Shopify/sarama"
	log "github.com/Sirupsen/logrus"
	"github.com/wvanbergen/kafka/consumergroup"
)

// MessageCallback is a short notation of a callback function for incoming Kafka message.
type MessageCallback func(msg *sarama.ConsumerMessage) error

// Consumer consumer messages.
type Consumer struct {
	ConsumerGroupName string
	ZkAddress         string
	Topic             string
}

type topic struct {
	Name            string
	MessageCallback MessageCallback
	Consuming       bool
}

// NewConsumer creates a new consumer.
func NewConsumer(consumerGroupName, topic, zkAddress string) (Consumer, error) {
	c := Consumer{}

	if consumerGroupName == "" {
		return c, fmt.Errorf(" consumer group name cannot be empty")
	}

	c = Consumer{
		ConsumerGroupName: consumerGroupName,
		ZkAddress:         zkAddress,
		Topic:             topic,
	}

	return c, nil
}

type messageProvider interface {
	Messages() <-chan *sarama.ConsumerMessage
	CommitUpto(msg *sarama.ConsumerMessage) error
}

// Start runs the process of consuming. It is blocking.
func (c *Consumer) Start() error {
	cg, err := consumergroup.JoinConsumerGroup(c.ConsumerGroupName, []string{c.Topic}, []string{c.ZkAddress}, nil)
	if err != nil {
		return err
	}
	defer cg.Close()

	runConsumer(c.Topic, cg)
	return nil
}

func runConsumer(topic string, provider messageProvider) {
	// Trap SIGINT to trigger a shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	for {
		select {
		case msg := <-provider.Messages():
			if msg.Topic != topic {
				log.Warnf("[%s] Cannot consume message. Handler not found.", msg.Topic)
				continue
			}

			log.Printf("[%s] Received message: %s", msg.Topic, string(msg.Value))

			if err := provider.CommitUpto(msg); err != nil {
				log.Warnf("[%s] Consuming message: %v", msg.Topic, err)
			}

			if string(msg.Value) == "freeze" {
				log.Warnf("[%s] This is one hard task!", msg.Topic)
				time.Sleep(10 * time.Second)
				log.Warnf("[%s] Finally done!", msg.Topic)
			}
		case <-signals:
			return
		}
	}
}
