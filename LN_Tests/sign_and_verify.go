package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"math/big"

	"fmt"
	"log"
)


// TrimmedCopy creates a trimmed copy of Transaction to be used in signing
func (tx *Transaction) TrimmedCopy() Transaction {
	var inputs []TXInput
	var outputs []TXOutput

	for _, vin := range tx.Vin {
		inputs = append(inputs, TXInput{vin.Txid, vin.Vout, nil, nil})
	}

	for _, vout := range tx.Vout {
		outputs = append(outputs, TXOutput{0, vout.Value, vout.PubKeyHash, timeLock{0, nil,nil}})
	}

	txCopy := Transaction{tx.ID, inputs, outputs}

	return txCopy
}

func (tx *Transaction) Sign(privKey ecdsa.PrivateKey) {
	//prevTXs is a map[TxId]Transactions
    if tx.IsCoinbase() {
		return
	}
    


	txCopy := tx.TrimmedCopy()
    //TrimmedCopy returns the same Tx with TXInputs without Signature and PubKey

	for inID, _ := range txCopy.Vin {
		txCopy.Vin[inID].Signature = nil
		txCopy.ID = txCopy.Hash()


		r, s, err := ecdsa.Sign(rand.Reader, &privKey, txCopy.ID)
		if err != nil {
			log.Panic(err)
		}
		signature := append(r.Bytes(), s.Bytes()...)

		tx.Vin[inID].Signature = append(tx.Vin[inID].Signature, signature...)

	}
}

func (tx *Transaction) SignSingleTXI(Idx int, privKey ecdsa.PrivateKey) {
	//prevTXs is a map[TxId]Transactions
    if tx.IsCoinbase() {
		return
	}


	txCopy := tx.TrimmedCopy()
    //TrimmedCopy returns the same Tx with TXInputs without Signature and PubKey

	txCopy.ID = txCopy.Hash()

    r, s, err := ecdsa.Sign(rand.Reader, &privKey, txCopy.ID)
	if err != nil {
		log.Panic(err)
	}
	signature := append(r.Bytes(), s.Bytes()...)


    tx.Vin[Idx].Signature = append(tx.Vin[Idx].Signature, signature...)


}


func (tx *Transaction) Verify(prevTXOs []TXOutput, currentHeight int64) bool {
	if tx.IsCoinbase() {
		return true
	}
    
        

	txCopy := tx.TrimmedCopy()


    verifying:

	for inIdx, vin := range tx.Vin {
        

        if vin.IsFund() {

            if !prevTXOs[inIdx].IsLockedWithKeys(vin.PubKey) {

                return false

            }

            if !txCopy.VerifyFundSig(vin) {
                return false
            }
            continue verifying
        }


        // check vin.pubKey with previous vout.pubKeyHash
        if !prevTXOs[inIdx].IsLockedWithKey(vin.PubKey[:keyLen]) {
            return false
        }



        if !vin.TimeLockExpired(prevTXOs[inIdx], currentHeight) {
            
            if !prevTXOs[inIdx].ValidTimeKey(vin.PubKey) {
            fmt.Println("Referenced TXO is time-locked, provided keys are invalid!")
            return false
            }


            if !txCopy.VerifyCommitSig(vin) {
                return false
            }
            continue verifying
        }


        if !txCopy.VerifyNormalSig(vin) {

            return false
        }

	}

	return true
}

func (txCopy *Transaction) VerifyNormalSig(vin TXInput) bool {


    curve := elliptic.P256()
    txCopy.ID = txCopy.Hash()

    singleKey := vin.PubKey[:keyLen]


	r := big.Int{}
	s := big.Int{}
	sigLen := len(vin.Signature)
	r.SetBytes(vin.Signature[:(sigLen / 2)])
	s.SetBytes(vin.Signature[(sigLen / 2):])

	x := big.Int{}
	y := big.Int{}
    x.SetBytes(singleKey[:(keyLen / 2)])
    y.SetBytes(singleKey[(keyLen / 2):])

	rawPubKey := ecdsa.PublicKey{curve, &x, &y}
    if ecdsa.Verify(&rawPubKey, txCopy.ID, &r, &s) == false {
        return false
    }
    return true
}

func (txCopy *Transaction) VerifyFundSig(vin TXInput) bool {

    pubKeyA := vin.PubKey[:keyLen]
    pubKeyB := vin.PubKey[keyLen:]

    curve := elliptic.P256()

	txCopy.ID = txCopy.Hash()


	r := big.Int{}
	s := big.Int{}
	sigLen := len(vin.Signature)
	r.SetBytes(vin.Signature[:(sigLen / 4)])
	s.SetBytes(vin.Signature[(sigLen / 4):(sigLen / 2)])

	x := big.Int{}
	y := big.Int{}
	x.SetBytes(pubKeyA[:(keyLen / 2)])
	y.SetBytes(pubKeyA[(keyLen / 2):])

	rawPubKey := ecdsa.PublicKey{curve, &x, &y}
	if ecdsa.Verify(&rawPubKey, txCopy.ID, &r, &s) == false {
		fmt.Println("FALSE_1")
        return false
	}

    R := big.Int{}
    S := big.Int{}
	R.SetBytes(vin.Signature[(sigLen / 2):(3*sigLen / 4)])
	S.SetBytes(vin.Signature[(3*sigLen / 4):])

	X := big.Int{}
	Y := big.Int{}
	X.SetBytes(pubKeyB[:(keyLen / 2)])
	Y.SetBytes(pubKeyB[(keyLen / 2):])

	rawPubKey = ecdsa.PublicKey{curve, &X, &Y}
	if ecdsa.Verify(&rawPubKey, txCopy.ID, &R, &S) == false {
        fmt.Println("FALSE_2")
		return false
	}

    return true
}


func (txCopy *Transaction) VerifyCommitSig(vin TXInput) bool {

    
    spPubKeyA := vin.PubKey[keyLen:(2 * keyLen)]
    spPubKeyB := vin.PubKey[(2 * keyLen):]

    curve := elliptic.P256()

	txCopy.ID = txCopy.Hash()

	r := big.Int{}
	s := big.Int{}
	sigLen := len(vin.Signature)
	r.SetBytes(vin.Signature[:(sigLen / 4)])
	s.SetBytes(vin.Signature[(sigLen / 4):(sigLen / 2)])

	x := big.Int{}
	y := big.Int{}
	x.SetBytes(spPubKeyA[:(keyLen / 2)])
	y.SetBytes(spPubKeyA[(keyLen / 2):])

	rawPubKey := ecdsa.PublicKey{curve, &x, &y}
	if ecdsa.Verify(&rawPubKey, txCopy.ID, &r, &s) == false {
		fmt.Println("FALSE_1")
        return false
	}

    R := big.Int{}
    S := big.Int{}
	R.SetBytes(vin.Signature[(keyLen / 2):(3*sigLen / 4)])
	S.SetBytes(vin.Signature[(3*sigLen / 4):])

	X := big.Int{}
	Y := big.Int{}
	X.SetBytes(spPubKeyB[:(keyLen / 2)])
	Y.SetBytes(spPubKeyB[(keyLen / 2):])

	rawPubKey = ecdsa.PublicKey{curve, &X, &Y}
	if ecdsa.Verify(&rawPubKey, txCopy.ID, &R, &S) == false {
        fmt.Println("FALSE_2")
		return false
	}

    
    return true
}

