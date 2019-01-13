package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"log"
)


// Wallet stores a serises of private and public keys and commitment
// transactions
type LitWallet struct {
    Opened          bool
    Party           bool
	PrivateKeys []ecdsa.PrivateKey
	PublicKeys  [][]byte
    OtherPrivKeys []ecdsa.PrivateKey
    Commitments     []*Transaction
}

// NewWallet creates and returns a Wallet
func NewLitWallet(party bool, zPrivateKey ecdsa.PrivateKey, zPublicKey []byte) *LitWallet {
    privateKeys := []ecdsa.PrivateKey{zPrivateKey}
    publicKeys := [][]byte{zPublicKey}
    litWallet := LitWallet{true, party, privateKeys, publicKeys, nil, nil}

	return &litWallet
}

func (lw *LitWallet) Close() {
    lw.Opened = false
}

func (lw *LitWallet) addKeyPair() (ecdsa.PrivateKey, []byte) {
	curve := elliptic.P256()
	private, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Panic(err)
	}
	pubKey := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)
    
    lw.PrivateKeys = append(lw.PrivateKeys, *private)
    lw.PublicKeys = append(lw.PublicKeys, pubKey)

	return *private, pubKey
}

func exchangeLatestKeys(litWalletA, litWalletB *LitWallet) {
    if len(litWalletA.Commitments) != len(litWalletB.Commitments) {
        log.Panic("ERROR: two parties have different number of commitments")
    }
    


    litWalletA.OtherPrivKeys = append(litWalletA.OtherPrivKeys, litWalletB.PrivateKeys[len(litWalletA.Commitments)])
    litWalletB.OtherPrivKeys = append(litWalletB.OtherPrivKeys, litWalletA.PrivateKeys[len(litWalletA.Commitments)])

}


func (lw *LitWallet) AddCommitment(commitment *Transaction) {
    lw.Commitments = append(lw.Commitments, commitment)
}

func (lw *LitWallet) GetKeyPair(version int) (ecdsa.PrivateKey, []byte) {
    return lw.PrivateKeys[version], lw.PublicKeys[version]
}


