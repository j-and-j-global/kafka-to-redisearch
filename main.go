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

	inputChan := make(chan MessageWithEnvelope, 1024)
	outputChan := make(chan MessageWithEnvelope, 1024)

	go func() {
		err := c.Kafka.ConsumerLoop(inputChan)
		if err != nil {
			panic(err)
		}
	}()

	go func() {
		err := c.Redis.WriteLoop(outputChan)
		if err != nil {
			panic(err)
		}
	}()

	panic(provenanceParserMap.TransformerLoop(inputChan, outputChan))
}
