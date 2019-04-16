package main

import (
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type kafkaClient interface {
	ReadMessage(time.Duration) (*kafka.Message, error)
	SubscribeTopics([]string, kafka.RebalanceCb) error
}

type Kafka struct {
	consumer kafkaClient
}

func NewKafka(bootstrapServers, topic string) (k Kafka, err error) {
	k.consumer, err = kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": bootstrapServers,
		"auto.offset.reset": "earliest",
		"group.id":          "kafka-to-redisearch",
	})

	if err != nil {
		return
	}

	err = k.consumer.SubscribeTopics([]string{topic}, nil)

	return
}

func (k Kafka) ConsumerLoop(c chan []byte) (err error) {
	var msg *kafka.Message

	for {
		msg, err = k.consumer.ReadMessage(-1)
		if err != nil {
			return
		}

		c <- msg.Value
	}

	return
}
