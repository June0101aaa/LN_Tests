package main


import (
//	"bytes"
//	"crypto/ecdsa"
//    "crypto/elliptic"
//	"crypto/rand"
//	"crypto/sha256"
//	"math/big"

//	"encoding/gob"
	"encoding/hex"
//	"fmt"
	"log"
//	"strings"
)

func NewFundingTx (alice, bob string, amountA, amountB int, UTXOSet *UTXOSet) *Transaction{
    var inputs []TXInput
    var outputs []TXOutput

    wallets, err := NewWallets()
	if err != nil {
		log.Panic(err)
	}
    walletA := wallets.GetWallet(alice)
    walletB := wallets.GetWallet(bob)

    fund := GetFundAddress(walletA.PublicKey, walletB.PublicKey)
    _, fundExist := wallets.Wallets[fund]
    if !fundExist {
        fund = wallets.CreateFundWallet(alice, bob)
    }

    wallets.SaveToFile()
    wallets, err = NewWallets()

	accA, validOutputsA := UTXOSet.FindSpendableOutputs(walletA.PublicKey, amountA)

    if accA < amountA {
		log.Panic("ERROR: %s Not enough funds\n", alice)
	}

    for txid, outs := range validOutputsA {
		txID, err := hex.DecodeString(txid)
		if err != nil {
			log.Panic(err)
		}

		for _, out := range outs {
			input := TXInput{txID, out, nil, walletA.PublicKey}
			inputs = append(inputs, input)
		}
	}


	accB, validOutputsB := UTXOSet.FindSpendableOutputs(walletB.PublicKey, amountB)

    if accB < amountB {
		log.Panic("ERROR: %s Not enough funds\n", bob)
	}

    for txid, outs := range validOutputsB {
		txID, err := hex.DecodeString(txid)
		if err != nil {
			log.Panic(err)
		}

		for _, out := range outs {
			input := TXInput{txID, out, nil, walletB.PublicKey}
			inputs = append(inputs, input)
		}
	}
    
    outputs = append(outputs, *NewTXOutput(amountA + amountB, fund))
    
    if accA > amountA {
		outputs = append(outputs, *NewTXOutput(accA-amountA, alice)) // a change
	}
    if accB > amountB {
		outputs = append(outputs, *NewTXOutput(accB-amountB, bob)) // a change
	}

	tx := Transaction{nil, inputs, outputs}
	tx.ID = tx.Hash()

    tx.ID = append(tx.ID, []byte("fund")...)
/*
    //debug
    fmt.Println(len(tx.ID))
*/


    for i := 0; i < len(validOutputsA); i++ {
        tx.SignSingleTXI(i, walletA.PrivateKey)
    }
    for i := 0; i < len(validOutputsB); i++ {
        tx.SignSingleTXI(len(validOutputsA) + i, walletB.PrivateKey)
    }
    
    return &tx
}


func GetFundAddress(alicePubKey, bobPubKey []byte) (string) {
    
    fundPublicKey := append(alicePubKey, bobPubKey...)
	pubKeyHash := HashPubKey(fundPublicKey)

	versionedPayload := append([]byte{version}, pubKeyHash...)
	checksum := checksum(versionedPayload)

	fullPayload := append(versionedPayload, checksum...)
	address := Base58Encode(fullPayload)

	return string(address)

}


