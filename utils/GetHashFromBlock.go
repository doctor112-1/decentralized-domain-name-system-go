package utils

import (
	"bytes"
	"encoding/gob"
	"log"
)

type GetHashFromBlock struct {
	GetHashFromBlock string
}

func (g *GetHashFromBlock) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(g)
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}

func DeserializeGetHashFromBlock(d []byte) *GetHashFromBlock {
	var GetHashFromBlockMessage GetHashFromBlock

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&GetHashFromBlockMessage)
	if err != nil {
		log.Panic(err)
	}

	return &GetHashFromBlockMessage
}
