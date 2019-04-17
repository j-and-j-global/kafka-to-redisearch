package main

import (
	"log"
)

func main() {
	defer func() {
		log.Print("Bye-bye :(")
	}()

	log.Print("Starting up!")

	c, err := NewConfig()
	if err != nil {
		panic(err)
	}

	log.Printf("Config: %+v", c)

	messageChan := make(chan []byte, 1024)

	go func() {
		err := c.Kafka.ConsumerLoop(messageChan)
		if err != nil {
			panic(err)
		}
	}()

	panic(c.Redis.WriteLoop(messageChan))
}
