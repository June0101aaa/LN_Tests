package main

import (
	"bytes"
	"crypto/elliptic"
	"encoding/gob"
//	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const litWalletFile = "lit_wallet.dat"

// Wallets stores a collection of wallets
type LitWallets struct {
	LitWallets map[string]*LitWallet
}

// NewWallets creates Wallets and fills it from a file if it exists
func NewLitWallets() (*LitWallets, error) {
    litWallets := LitWallets{}
	litWallets.LitWallets = make(map[string]*LitWallet)

	err := litWallets.LoadFromFile()

	return &litWallets, err
}


// CreateWallet adds a Wallet to Wallets
func (lws *LitWallets) CreateLitWallet(party bool, wallet Wallet) {
	litWallet := NewLitWallet(party, wallet.PrivateKey, wallet.PublicKey)

    lws.LitWallets[wallet.GetAddress()] = litWallet
}

// AddCommitment adds a new version of commitmentTx
func (lws *LitWallets) NewCommitment(address string, commitTx *Transaction) {
    lws.LitWallets[address].AddCommitment(commitTx)


}


/*
// GetAddresses returns an array of addresses stored in the wallet file
func (ws *Wallets) GetAddresses() []string {
	var addresses []string

	for address := range ws.Wallets {
		addresses = append(addresses, address)
	}

	return addresses
}
*/

// GetWallet returns a Wallet by its address
func (lws LitWallets) GetWallet(address string) *LitWallet {
	return lws.LitWallets[address]
}


// LoadFromFile loads wallets from the file
func (lws *LitWallets) LoadFromFile() error {
	if _, err := os.Stat(litWalletFile); os.IsNotExist(err) {
		return err
	}

	fileContent, err := ioutil.ReadFile(litWalletFile)
	if err != nil {
		log.Panic(err)
	}

	var litWallets LitWallets
	gob.Register(elliptic.P256())
	decoder := gob.NewDecoder(bytes.NewReader(fileContent))
	err = decoder.Decode(&litWallets)
	if err != nil {
		log.Panic(err)
	}

	lws.LitWallets = litWallets.LitWallets

	return nil
}

// SaveToFile saves wallets to a file
func (lws LitWallets) SaveToFile() {
	var content bytes.Buffer

	gob.Register(elliptic.P256())

	encoder := gob.NewEncoder(&content)
	err := encoder.Encode(lws)
	if err != nil {
		log.Panic(err)
	}

	err = ioutil.WriteFile(litWalletFile, content.Bytes(), 0644)
	if err != nil {
		log.Panic(err)
	}
}
