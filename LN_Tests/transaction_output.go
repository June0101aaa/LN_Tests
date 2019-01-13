package main

import (
	"bytes"
	"encoding/gob"
	"log"
    "fmt"
)

// TXOutput represents a transaction output
type TXOutput struct {
    Height     int64
	Value      int
	PubKeyHash []byte
    TimeLock   timeLock
}



// IsLockedWithKey checks if the output can be used by the owner of the pubkey
func (out *TXOutput) IsLockedWithKey(pubKey []byte) bool {
	return bytes.Compare(out.PubKeyHash, HashPubKey(pubKey)) == 0
}

func (out *TXOutput) IsLockedWithKeys(pubKey []byte) bool {
    if len(pubKey) < 2 * keyLen {
        fmt.Println("Keys too short")
        return false
    }
	return bytes.Compare(out.PubKeyHash, HashPubKey(pubKey)) == 0
}


// NewTXOutput create a new TXOutput
func NewTXOutput(value int, address string) *TXOutput {
	txo := &TXOutput{0, value, nil, timeLock{0, nil, nil}}
	txo.Lock([]byte(address))

	return txo
}



// TXOutputs collects TXOutput
type TXOutputs struct {
	Outputs []TXOutput
}

// Serialize serializes TXOutputs
func (outs TXOutputs) Serialize() []byte {
	var buff bytes.Buffer

	enc := gob.NewEncoder(&buff)
	err := enc.Encode(outs)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}

// DeserializeOutputs deserializes TXOutputs
func DeserializeOutputs(data []byte) TXOutputs {
	var outputs TXOutputs

	dec := gob.NewDecoder(bytes.NewReader(data))
	err := dec.Decode(&outputs)
	if err != nil {
		log.Panic(err)
	}

	return outputs
}
