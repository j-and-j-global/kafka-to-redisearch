package main

func main() {
	c, err := NewConfig()
	if err != nil {
		panic(err)
	}

	messageChan := make(chan []byte, 1024)

	go func() {
		err := c.Kafka.ConsumerLoop(messageChan)
		if err != nil {
			panic(err)
		}
	}()

	panic(c.Redis.WriteLoop(messageChan))
}
