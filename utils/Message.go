package utils

import (
	"bytes"
	"encoding/gob"
	"log"
	"unsafe"
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

func CreateMessage(command string, serializedPayload []byte) *Message {
	var Message Message

	Message.Version = 1
	Message.Command = command
	Message.Payload = string(serializedPayload)
	Message.Length = uint32(unsafe.Sizeof(Message.Version)) + uint32(unsafe.Sizeof(Message.Command)) + uint32(unsafe.Sizeof(Message.Payload))

	return &Message
}
