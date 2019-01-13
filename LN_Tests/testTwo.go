package main

import (
    "fmt"
//    "log"
    "time"
)


func testTwo() {
    cli := CLI{}

    wallets, _ := NewWallets()
    alice := wallets.CreateWallet()
    bob :=  wallets.CreateWallet()
    miner :=  wallets.CreateWallet()
    wallets.SaveToFile()

    cli.createBlockchain(miner)

    for i := 0; i < 5; i++ {
        cli.send(miner, alice, miner, 20)
        cli.send(miner, bob, miner, 20)
    }

    fmt.Println("Alice getBalance - ")
    cli.getBalance(alice)   
    fmt.Println("Bob getBalance - ")
    cli.getBalance(bob)

    fmt.Println("======================Start of 30 transactions======================") 
    start := time.Now()

    cli.newFund(alice, bob, miner, 50, 50)

    for i := 0; i < 30; i++ {
        fmt.Printf("Lightning Transaction %v: ", i + 1)
        cli.litSend(alice, bob, 1)
    }

    cli.endFund(alice, bob, miner)

    end := time.Now()
    elapsed := end.Sub(start)

    fmt.Println("======================End of 30 transactions======================") 

    fmt.Println("Alice getBalance - ")
    cli.getBalance(alice)   
    fmt.Println("Bob getBalance - ")
    cli.getBalance(bob)

    fmt.Println("30 transactions through Lightning Network - execution time: ", elapsed)

}

