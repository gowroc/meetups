package kafka

import (
	"encoding/json"
	"fmt"

	"github.com/Shopify/sarama"
)

// Producer is interface for sending messages to Kafka.
type Producer interface {
	SendMessage(topic string, message interface{}) error
}

type producer struct {
	Producer sarama.SyncProducer
}

// NewProducer returns a new SyncProducer for give brokers addresses.
func NewProducer(broker string) (Producer, error) {
	kafkaConfig := sarama.NewConfig()
	kafkaConfig.Producer.Return.Successes = true

	kafkaProducer, err := sarama.NewSyncProducer([]string{broker}, kafkaConfig)
	if err != nil {
		return nil, err
	}

	return &producer{Producer: kafkaProducer}, nil
}

func (p *producer) SendMessage(topic string, message interface{}) error {
	msg, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("cannot marshal message %v: %v", message, err)
	}

	_, _, err = p.Producer.SendMessage(&sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(msg),
	})
	if err != nil {
		return fmt.Errorf("cannot send message %v: %v", message, err)
	}

	return nil
}
