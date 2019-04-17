package main

import (
	"log"

	"github.com/RediSearch/redisearch-go/redisearch"
)

type redisClient interface {
	CreateIndex(*redisearch.Schema) error
	IndexOptions(redisearch.IndexingOptions, ...redisearch.Document) error
	Info() (*redisearch.IndexInfo, error)
}

type Redis struct {
	client redisClient
}

func NewRedis(master, index string) (r Redis, err error) {
	r.client = redisearch.NewClient(master, index)

	info, _ := r.client.Info()
	if info != nil {
		return
	}

	sc := redisearch.NewSchema(redisearch.DefaultOptions).
		AddField(redisearch.NewTextField("body")).
		AddField(redisearch.NewTextField("author")).
		AddField(redisearch.NewTextFieldOptions("title", redisearch.TextFieldOptions{Weight: 5.0, Sortable: true})).
		AddField(redisearch.NewNumericField("date"))

	err = r.client.CreateIndex(sc)

	return
}

func (r Redis) WriteLoop(c chan []byte) (err error) {
	var message MessageWithEnvelope

	indexOpts := redisearch.DefaultIndexingOptions
	indexOpts.Replace = true

	for m := range c {
		log.Printf("Redis Loop: Handling %q", string(m))

		message, err = ParseMessage(m)
		if err != nil {
			log.Printf("Redis Loop: invalid message: %+v", err)

			continue
		}

		if message.Create() || message.Update() {
			doc := redisearch.NewDocument(message.Message.Slug, 1.0)
			doc.Set("body", message.Message.Body).
				Set("author", message.Message.Author).
				Set("title", message.Message.Title).
				Set("date", message.Message.Date.Unix())

			err = r.client.IndexOptions(indexOpts, []redisearch.Document{doc}...)
			if err != nil {
				return
			}
		} else if message.Delete() {
			log.Print("Redis Loop: deletion NOP")
		}
	}

	return
}
