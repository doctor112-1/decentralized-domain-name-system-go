package utils

import (
	"bytes"
	"encoding/gob"
	"log"
)

type SendIP struct {
	SendIP string
}

func (s *SendIP) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(s)
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}

func DeserializeSendIP(d []byte) *SendIP {
	var SendIPMessage SendIP

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&SendIPMessage)
	if err != nil {
		log.Panic(err)
	}

	return &SendIPMessage
}
