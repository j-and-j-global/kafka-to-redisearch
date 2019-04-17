package main

import (
	"log"

	"github.com/RediSearch/redisearch-go/redisearch"
)

type redisClient interface {
	CreateIndex(*redisearch.Schema) error
	Index(...redisearch.Document) error
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
	var message Message

	for m := range c {
		message, err = NewMessage(m)
		if err != nil {
			log.Print("%+v is an invalid message: %+v", m, err)

			continue
		}

		doc := redisearch.NewDocument(message.Slug, 1.0)
		doc.Set("body", message.Body).
			Set("author", message.Author).
			Set("title", message.Title).
			Set("date", message.Date)

		err = r.client.Index([]redisearch.Document{doc}...)
		if err != nil {
			return
		}
	}

	return
}
