package main

import (
	"bytes"
	"encoding/gob"
	"log"
)

type InvalidBlock struct {
	InvalidBlock string
}

func (i *InvalidBlock) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(i)
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}

func DeserializeInvalidBlock(d []byte) *InvalidBlock {
	var InvalidBlockMessage InvalidBlock

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&InvalidBlockMessage)
	if err != nil {
		log.Panic(err)
	}

	return &InvalidBlockMessage
}
