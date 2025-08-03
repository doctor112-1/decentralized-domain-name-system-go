package main

import (
	"bytes"
	"encoding/gob"
	"log"
)

type ConnectNodeResponse struct {
	ConnectNode  uint8
	BlocksAmount uint64
}

func (c *ConnectNodeResponse) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(c)
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}

func DeserializeConnectNodeResponse(d []byte) *ConnectNodeResponse {
	var ConnectNodeResponseMessage ConnectNodeResponse

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&ConnectNodeResponseMessage)
	if err != nil {
		log.Panic(err)
	}

	return &ConnectNodeResponseMessage
}
