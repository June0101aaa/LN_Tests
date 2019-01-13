package main

import (
	"bytes"
    "fmt"
)


type timeLock struct {
    Time        int
    PubKeyHashA []byte
    PubKeyHashB []byte
}


// Lock signs the output
func (out *TXOutput) Lock(address []byte) {
	pubKeyHash := Base58Decode(address)
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-4]
	out.PubKeyHash = pubKeyHash
}

func (out *TXOutput) ValidTimeKey(pubKey []byte) bool {

    if len(pubKey) < 3 * keyLen {
        fmt.Println("Keys too short")
        return false
    }
    pubKeyA := pubKey[keyLen:(2 * keyLen)]
    pubKeyB := pubKey[(2 * keyLen):]

    return bytes.Compare(out.TimeLock.PubKeyHashA, HashPubKey(pubKeyA)) == 0 && 
        bytes.Compare(out.TimeLock.PubKeyHashB, HashPubKey(pubKeyB)) == 0
}

func (in *TXInput) TimeLockExpired(prevTXO TXOutput, currentHeight int64) bool {
    prevHeight := prevTXO.Height


    if int(currentHeight - prevHeight) < prevTXO.TimeLock.Time {
            return false
    }

    return true
}

