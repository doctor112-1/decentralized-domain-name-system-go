package utils

import (
	"bytes"
	"encoding/gob"
	"log"
)

type SendNodes struct {
	SendNodes string
}

func (s *SendNodes) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(s)
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}

func DeserializeSendNodes(d []byte) *SendNodes {
	var SendNodesMessage SendNodes

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&SendNodesMessage)
	if err != nil {
		log.Panic(err)
	}

	return &SendNodesMessage
}
