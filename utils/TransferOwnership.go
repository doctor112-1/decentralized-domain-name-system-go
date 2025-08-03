package utils

import (
	"bytes"
	"encoding/gob"
	"log"
)

type TransferOwnership struct {
	TransferOwndership string
}

func (t *TransferOwnership) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(t)
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}

func DeserializeTransferOwnership(d []byte) *TransferOwnership {
	var TransferOwnershipMessage TransferOwnership

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&TransferOwnershipMessage)
	if err != nil {
		log.Panic(err)
	}

	return &TransferOwnershipMessage
}
