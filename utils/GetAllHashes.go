package main

import (
	"bytes"
	"encoding/gob"
	"log"
)

type GetAllHashes struct {
	GetAllHashes uint8
}

func (g *GetAllHashes) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(g)
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}

func DeserializeGetAllHashes(d []byte) *GetAllHashes {
	var GetAllHashesMessage GetAllHashes

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&GetAllHashesMessage)
	if err != nil {
		log.Panic(err)
	}

	return &GetAllHashesMessage
}
