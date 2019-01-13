package main

import (
    "fmt"
    "log"

)

func (cli *CLI) newFund(alice, bob, miner string, amountA, amountB int) {
	if !ValidateAddress(alice) {
		log.Panic("ERROR: ADDRESS1 address is not valid")
	}
	if !ValidateAddress(bob) {
		log.Panic("ERROR: ADDRESS2 address is not valid")
	}

    wallets, _ := NewWallets()
 

 
 
    // 1. Create FundingTx
	bc := NewBlockchain()
	UTXOSet := UTXOSet{bc}
	defer bc.db.Close()

    fundTx := NewFundingTx(alice, bob, amountA, amountB, &UTXOSet)

    // 2. create LitWallets, LitWallet
    litWallets, _ := NewLitWallets()
    litWallets.CreateLitWallet(false, wallets.GetWallet(alice))
    litWallets.CreateLitWallet(true, wallets.GetWallet(bob))
    litWallets.SaveToFile()
    
    // 3. Create 1st CommitmentTx
    FirstCommitmentTx(alice, bob, fundTx, amountA, amountB)


    // 4. mine FundingTx
    cbTx := NewCoinbaseTX(miner, "")
	txs := []*Transaction{cbTx, fundTx}

	newBlock := bc.MineBlock(txs)
	UTXOSet.Update(newBlock)
	fmt.Println("Success!")

    fmt.Println("FundingTx created: ", alice, " amount - ", amountA, " ", bob,  " amount - ", amountB)
    fmt.Println("FundingTx broadcasted on to blockchain!")


}


