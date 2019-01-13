package main

import (
    "fmt"
    "time"
)


func testOne() {
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


    for i := 0; i < 30; i++ {
        fmt.Printf("Bitcoin Transaction %v: \n", i + 1)
        cli.send(alice, bob, miner, 1)
    }


    end := time.Now()
    elapsed := end.Sub(start)
    
    fmt.Println("======================End of 30 transactions======================")

    fmt.Println("Alice getBalance - ")
    cli.getBalance(alice)   
    fmt.Println("Bob getBalance - ")
    cli.getBalance(bob)

    fmt.Println("30 transactions on blockchain - execution time: ", elapsed)

}

