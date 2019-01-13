package main

import (
    "bytes"
)

// TXInput represents a transaction input
type TXInput struct {
	Txid      []byte
	Vout      int
	Signature []byte
	PubKey    []byte
}

// UsesKey checks whether the address initiated the transaction
func (in *TXInput) UsesKey(pubKeyHash []byte) bool {
	lockingHash := HashPubKey(in.PubKey)

	return bytes.Compare(lockingHash, pubKeyHash) == 0
}

func (in *TXInput) IsFund() bool {

    return bytes.Compare(in.Txid[32:], []byte("fund")) == 0 && len(in.PubKey) == 2 * keyLen
}



