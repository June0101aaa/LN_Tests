package main

import (
    "fmt"
    "log"
)

func testFour() {

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
    
    fmt.Println("Alice getBalance - ")
    cli.getBalance(alice)   
    fmt.Println("Bob getBalance - ")
    cli.getBalance(bob)

    cli.newFund(alice, bob, miner, 6, 6)

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







    for i := 0; i < lockTime + 1 ; i++ {
        cli.send(miner, chris, miner, 5)
    }
    fmt.Println("=============================================================")      
    fmt.Println("Alice(", alice, ") spends money after lockTime + 1 = ",lockTime + 1 , " transactions")


    cli.send(alice, chris, miner,balance)


}
