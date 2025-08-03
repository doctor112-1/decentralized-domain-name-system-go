package main

import (
	"bytes"
	"encoding/gob"
	"log"
)

type Block struct {
	Timestamp  uint64
	Nonce      uint32
	Target     uint32
	DomainName string
	Hash       string
	HashRecord string
	PublicKey  string
}

func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(b)
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}

func DeserializeBlock(d []byte) *Block {
	var Blockblock Block

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&Blockblock)
	if err != nil {
		log.Panic(err)
	}

	return &Blockblock
}
