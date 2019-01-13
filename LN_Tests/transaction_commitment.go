package main

import (

	"log"

)

const lockTime = 5


func FirstCommitmentTx(alice, bob string, fundTx *Transaction, amount1, amount2 int) {

    //
    wallets, errOne := NewWallets()
	if errOne != nil {
		log.Panic(errOne)
	}

    publicA := wallets.GetWallet(alice).PublicKey
    publicB := wallets.GetWallet(bob).PublicKey
    
    fund := GetFundAddress(publicA, publicB)
	fundWallet := wallets.GetWallet(fund)
    
    litWallets, errTwo := NewLitWallets()    
    if errTwo != nil {
		log.Panic(errTwo)
	}
    defer litWallets.SaveToFile()

    litWalletA := litWallets.GetWallet(alice)
    litWalletB := litWallets.GetWallet(bob)

    _, publicAa := litWalletA.addKeyPair()
    _, publicBa := litWalletB.addKeyPair()



  
    var inputsA []TXInput
    var inputsB []TXInput
    var outputsA []TXOutput
    var outputsB []TXOutput

    input := TXInput{fundTx.ID, 0, nil, fundWallet.PublicKey}
    inputsA = append(inputsA, input)
    inputsB = append(inputsB, input)

    outputsA = append(outputsA, *NewCommitTXOutput(amount1, alice, publicAa, publicB))
    outputsA = append(outputsA, *NewTXOutput(amount2, bob))

    outputsB = append(outputsB, *NewTXOutput(amount1, alice))
    outputsB = append(outputsB, *NewCommitTXOutput(amount2, bob, publicA, publicBa))

    txA := Transaction{nil, inputsA, outputsA}
	txA.ID = txA.Hash()

    txB := Transaction{nil, inputsB, outputsB}
	txB.ID = txB.Hash()
    
    txA.Sign(litWalletA.PrivateKeys[0])
    txB.Sign(litWalletB.PrivateKeys[0])

    litWallets.NewCommitment(alice, &txA)
    litWallets.NewCommitment(bob, &txB)
    
}

// NewCommitmentTx
func NewCommitmentTx(from, to string, amount int) {
    
    //Check for information
    aliceToBob := true

    litWallets, err := NewLitWallets()
    if err != nil {
		log.Panic(err)
	}
    defer litWallets.SaveToFile()

    litWalletA := litWallets.GetWallet(from)

    alice := from
    bob := to

    if litWalletA.Party {
        alice = to
        bob = from
        aliceToBob = false
    }  

    litWalletA = litWallets.GetWallet(alice)
    litWalletB := litWallets.GetWallet(bob)
    
    //TODO: Check if the network has been created

    version := len(litWalletA.Commitments)
    txA := *litWalletA.Commitments[version-1]
    txA.Vin[0].Signature = nil
    input := txA.Vin[0]
    aliceBalance := txA.Vout[0].Value
    bobBalance := txA.Vout[1].Value

    if (aliceToBob && amount > aliceBalance) || (!aliceToBob && amount > bobBalance) {
        log.Panic("ERROR: Not enough funds")
    }

    if aliceToBob {
        amount = 0 - amount
    }
    

    // Start a new commitment
    exchangeLatestKeys(litWalletA, litWalletB)

    publicA := litWalletA.PublicKeys[0]
    publicB := litWalletB.PublicKeys[0]

    _, publicAa := litWalletA.addKeyPair()
    _, publicBa := litWalletB.addKeyPair()

    var inputsA []TXInput
    var inputsB []TXInput
    var outputsA []TXOutput
    var outputsB []TXOutput

    inputsA = append(inputsA, input)
    inputsB = append(inputsB, input)
 
    outputsA = append(outputsA, *NewCommitTXOutput(aliceBalance + amount, alice, publicAa, publicB))
    outputsA = append(outputsA, *NewTXOutput(bobBalance - amount, bob))

    outputsB = append(outputsB, *NewTXOutput(aliceBalance + amount, alice))
    outputsB = append(outputsB, *NewCommitTXOutput(bobBalance - amount, bob, publicA, publicBa))



    txA = Transaction{nil, inputsA, outputsA}
	txA.ID = txA.Hash()

    txB := Transaction{nil, inputsB, outputsB}
	txB.ID = txB.Hash()
    
    txA.Sign(litWalletB.PrivateKeys[0])
    txB.Sign(litWalletA.PrivateKeys[0])
    

    litWallets.NewCommitment(alice, &txA)
    litWallets.NewCommitment(bob, &txB)


}

// NewCommitTXOutput
func NewCommitTXOutput(value int, address string, publicAa, publicBa []byte) *TXOutput {
	ctxo := &TXOutput{0, value, nil, timeLock{0, nil, nil}}
	ctxo.Lock([]byte(address))
    ctxo.SetTimeLock(publicAa, publicBa)

	return ctxo
}

// SetTimeLock
func (ctxo *TXOutput) SetTimeLock(publicAa, publicBa []byte) {
    ctxo.TimeLock = timeLock{lockTime, HashPubKey(publicAa), HashPubKey(publicBa)}
}

func SignCommitment(address string, version int) (*Transaction) {
    
    litWallets, err := NewLitWallets()
    if err != nil {
		log.Panic(err)
	}
    defer litWallets.SaveToFile()

    litWallet := litWallets.GetWallet(address)
    signDirectly := litWallet.Party

    if version > len(litWallet.Commitments) {
        log.Panic("ERROR: No such version!")
    }
    tx := litWallet.Commitments[len(litWallet.Commitments)-version]




    if !signDirectly {



        signature_1 := tx.Vin[0].Signature
        tx.Vin[0].Signature = nil
        tx.Sign(litWallet.PrivateKeys[0])
        tx.Vin[0].Signature = append(tx.Vin[0].Signature, signature_1 ...)


        return tx
    }

    tx.Sign(litWallet.PrivateKeys[0])

    return tx
}


