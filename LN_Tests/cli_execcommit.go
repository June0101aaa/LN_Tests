package main


import (
    "fmt"
    "log"
)

func (cli *CLI) execCommit(address, miner string, version int) {
    
    if !ValidateAddress(address) {
		log.Panic("ERROR: address is not valid")
	}
    
    bc := NewBlockchain()
	UTXOSet := UTXOSet{bc}
	defer bc.db.Close()

    commitmentTx := SignCommitment(address, version)

    cbTx := NewCoinbaseTX(miner, "")
	txs := []*Transaction{cbTx, commitmentTx}

    newBlock := bc.MineBlock(txs)
	UTXOSet.Update(newBlock)
	fmt.Println("Success!")


    
    fmt.Println("Commitment broadcasted!")

}

