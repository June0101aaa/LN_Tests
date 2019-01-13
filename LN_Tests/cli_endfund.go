package main

import (
    "fmt"
    "log"

)

func (cli *CLI) endFund(add_1, add_2, miner string) {
    
    if !ValidateAddress(add_1) {
		log.Panic("ERROR: ADDRESS1 address is not valid")
	}
	if !ValidateAddress(add_2) {
		log.Panic("ERROR: ADDRESS2 address is not valid")
	}
    
    bc := NewBlockchain()
	UTXOSet := UTXOSet{bc}
	defer bc.db.Close()


    settleTx := NewSettlementTx(add_1, add_2)

    cbTx := NewCoinbaseTX(miner, "")
	txs := []*Transaction{cbTx, settleTx}

    newBlock := bc.MineBlock(txs)
	UTXOSet.Update(newBlock)
	fmt.Println("Exercise Settlement Tx!")
    
    /*
    litWallets, err := NewLitWallets()
    defer litWallets.SaveToFile()
    if err != nil {
		log.Panic(err)
	}

    litWallet := litWallets.GetWallet(add_1)
    litWallet.Close()
    litWallet = litWallets.GetWallet(add_2)
    litWallet.Close()
    */

    
}

