package utils

import (
	"bytes"
	"encoding/gob"
	"log"
)

type Message struct {
	Length  uint32
	Version uint32
	Command string
	Payload string
}

func (m *Message) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(m)
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}

func DeserializeMessage(d []byte) *Message {
	var Messagemessage Message

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&Messagemessage)
	if err != nil {
		log.Panic(err)
	}

	return &Messagemessage
}
