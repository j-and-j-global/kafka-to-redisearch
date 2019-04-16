package main

import (
	"os"
)

type Config struct {
	Kafka Kafka
	Redis Redis
}

func NewConfig() (c Config, err error) {
	c.Kafka, err = NewKafka(os.Getenv("KAFKA_BOOTSTRAP"), os.Getenv("KAFKA_TOPIC"))
	if err != nil {
		return
	}

	c.Redis, err = NewRedis(os.Getenv("REDIS_MASTER"), os.Getenv("REDIS_INDEX"))

	return
}
