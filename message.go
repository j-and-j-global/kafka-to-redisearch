package main

import (
	"encoding/json"
	"time"
)

var (
	CreateOperation = "CREATE"
	UpdateOperation = "UPDATE"
	DeleteOperation = "DELETE"
)

type MessageWithEnvelope struct {
	Operation string
	Message   Message
}

func (m MessageWithEnvelope) Create() bool {
	return m.Operation == CreateOperation
}

func (m MessageWithEnvelope) Update() bool {
	return m.Operation == UpdateOperation
}

func (m MessageWithEnvelope) Delete() bool {
	return m.Operation == DeleteOperation
}

type Message struct {
	Slug   string
	Title  string
	Author string
	Date   time.Time
	Body   string
}

func ParseMessage(b []byte) (m MessageWithEnvelope, err error) {
	err = json.Unmarshal(b, &m)

	return
}
