package main

import (
	"fmt"
	"log"
)

func (cli *CLI) getBalance(address string) {
	if !ValidateAddress(address) {
		log.Panic("ERROR: Address is not valid")
	}
	bc := NewBlockchain()
	UTXOSet := UTXOSet{bc}
	defer bc.db.Close()

	balance := 0
	wallets, err := NewWallets()
	if err != nil {
		log.Panic(err)
	}
	wallet := wallets.GetWallet(address)
	UTXOs := UTXOSet.FindUTXO(wallet.PublicKey)

	for _, out := range UTXOs {
		balance += out.Value
	}

	fmt.Printf("Balance of '%s': %d\n", address, balance)
}
