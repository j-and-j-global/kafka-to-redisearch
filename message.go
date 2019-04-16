package main

import (
	"encoding/json"
	"time"
)

type Message struct {
	Slug   string
	Title  string
	Author string
	Date   time.Time
	Body   string
}

func NewMessage(b []byte) (m Message, err error) {
	err = json.Unmarshal(b, &m)

	return
}
