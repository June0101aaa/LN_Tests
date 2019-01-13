package main

import (
	"fmt"
	"log"

)

func (cli *CLI) send(from, to, miner string, amount int) {
	if !ValidateAddress(from) {
		log.Panic("ERROR: Sender address is not valid")
	}
	if !ValidateAddress(to) {
		log.Panic("ERROR: Recipient address is not valid")
	}

	bc := NewBlockchain()
	UTXOSet := UTXOSet{bc}
	defer bc.db.Close()


	tx := NewUTXOTransaction(from, to, amount, &UTXOSet)
	cbTx := NewCoinbaseTX(miner, "")
	txs := []*Transaction{cbTx, tx}



	newBlock := bc.MineBlock(txs)
	UTXOSet.Update(newBlock)
    fmt.Println(from, " sends ", to, ": ", amount, "coin(s)")
}
