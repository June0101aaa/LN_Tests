package main

import (
    "fmt"
    "log"
)

func testThree() {

    cli := CLI{}

    wallets, err := NewWallets()
    alice := wallets.CreateWallet()
    bob :=  wallets.CreateWallet()
	miner := wallets.CreateWallet()
    chris := wallets.CreateWallet()
	wallets.SaveToFile()

    cli.createBlockchain(miner)
    
    cli.send(miner, alice, miner, 10)
    cli.send(miner, bob, miner, 10)
    
    //Get Balance
    fmt.Println("Alice getBalance - ")
    cli.getBalance(alice)   
    fmt.Println("Bob getBalance - ")
    cli.getBalance(bob)

    cli.newFund(alice, bob, miner, 6, 6)
    
    //Get Balance
    fmt.Println("======================After Funding Tx======================")
    fmt.Println("Alice getBalance - ")
    cli.getBalance(alice)   
    fmt.Println("Bob getBalance - ")
    cli.getBalance(bob)


    cli.litSend(alice, bob, 2)
    cli.litSend(bob, alice, 1)
    cli.execCommit(alice, miner, 1)

    fmt.Println("======================Alice broadcasts Commitment Tx prematurely======================")


    //Get balance
    fmt.Println("Alice getBalance - ")
    cli.getBalance(alice)   
    fmt.Println("Bob getBalance - ")
    cli.getBalance(bob)

	bc := NewBlockchain()
	UTXOSet := UTXOSet{bc}

    balance := 0
	wallets, err = NewWallets()
	if err != nil {
		log.Panic(err)
	}
	wallet := wallets.GetWallet(alice)
	UTXOs := UTXOSet.FindUTXO(wallet.PublicKey)

	for _, out := range UTXOs {
		balance += out.Value
	}
    bc.db.Close()
    wallets.SaveToFile()

    fmt.Println("=============================================================")

    fmt.Println("Alice(", alice, ") wants to spend money immediately")

    cli.send(alice, chris, miner, balance)


}
