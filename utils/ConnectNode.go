package main

import (
	"bytes"
	"encoding/gob"
	"log"
)

type ConnectNode struct {
	ConnectNode  uint8
	BlocksAmount uint64
}

func (c *ConnectNode) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(c)
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}

func DeserializeConnectNode(d []byte) *ConnectNode {
	var ConnectNodeMessage ConnectNode

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&ConnectNodeMessage)
	if err != nil {
		log.Panic(err)
	}

	return &ConnectNodeMessage
}
