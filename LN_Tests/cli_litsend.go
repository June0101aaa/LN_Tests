package main

import (
    "fmt"
    "log"
)

func (cli *CLI) litSend(from, to string, amount int) {
    
    if !ValidateAddress(from) {
		log.Panic("ERROR: Sender address is not valid")
	}
	if !ValidateAddress(to) {
		log.Panic("ERROR: Recipient address is not valid")
	}
    
    NewCommitmentTx(from , to, amount)

    fmt.Println(from, " sends ", to, ": ", amount, "coin(s)")
}

