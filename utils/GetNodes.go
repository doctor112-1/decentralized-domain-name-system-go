package utils

import (
	"bytes"
	"encoding/gob"
	"log"
)

type GetNodes struct {
	GetNodes uint8
}

func (g *GetNodes) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(g)
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}

func DeserializeGetNodes(d []byte) *GetNodes {
	var GetNodesMessage GetNodes

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&GetNodesMessage)
	if err != nil {
		log.Panic(err)
	}

	return &GetNodesMessage
}
