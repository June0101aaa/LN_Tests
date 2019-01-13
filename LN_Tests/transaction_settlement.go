package main


import (
	"crypto/ecdsa"

	"log"
)

func NewSettlementTx (add_1, add_2 string) *Transaction{
    litWallets, err := NewLitWallets()
    defer litWallets.SaveToFile()
    if err != nil {
		log.Panic(err)
	}
    litWalletA := litWallets.GetWallet(add_1)
    
    alice := add_1
    bob := add_2

    if litWalletA.Party {
        alice = add_2
        bob = add_1
    }  

    litWalletA = litWallets.GetWallet(alice)
    litWalletB := litWallets.GetWallet(bob)
    
    //TODO: check if lightning network created

    privKeyA, pubKeyA := litWalletA.GetKeyPair(0)
    privKeyB, _ := litWalletB.GetKeyPair(0)

    txDraft := litWalletA.Commitments[len(litWalletA.Commitments)-1]
    txDraft.Vin[0].Signature = nil

    var inputs []TXInput
    var outputs []TXOutput

    inputs = append(inputs, txDraft.Vin[0])

    output_0 := txDraft.Vout[0]
    output_1 := txDraft.Vout[1]

    output_0.PubKeyHash = HashPubKey(pubKeyA)
    output_0.TimeLock = timeLock{0, nil, nil}

    outputs = append(outputs, []TXOutput{output_0, output_1}...)

    tx := Transaction{nil, inputs, outputs}
	tx.ID = tx.Hash()
    tx.SignSettlement(privKeyA, privKeyB)

    return &tx

}

func (tx *Transaction) SignSettlement(privKeyA, privKeyB ecdsa.PrivateKey) {
    tx.Sign(privKeyA)
    tx.Sign(privKeyB)
}
