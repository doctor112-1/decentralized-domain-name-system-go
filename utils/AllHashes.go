package utils

import (
	"bytes"
	"encoding/gob"
	"log"
)

type AllHashes struct {
	AllHashes string
}

func (a *AllHashes) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(a)
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}

func DeserializeAllHashes(d []byte) *AllHashes {
	var AllHashesMessage AllHashes

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&AllHashesMessage)
	if err != nil {
		log.Panic(err)
	}

	return &AllHashesMessage
}
