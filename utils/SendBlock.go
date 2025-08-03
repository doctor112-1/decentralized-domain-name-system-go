package main

import (
	"bytes"
	"encoding/gob"
	"log"
)

type SendBlock struct {
	SendBlock string
}

func (s *SendBlock) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(s)
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}

func DeserializeSendBlock(d []byte) *SendBlock {
	var SendBlockMessage SendBlock

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&SendBlockMessage)
	if err != nil {
		log.Panic(err)
	}

	return &SendBlockMessage
}
