package main

import (
	"crypto/ecdsa"
)

type block struct {
	index        int
	timestamp    string
	data         string
	transaction  []transaction
	Nonce        int
	previousHash string
	hash         string
}

type wallet struct {
	privateKey *ecdsa.PrivateKey
	publicKey  []byte
}

type transaction struct {
	sender    string
	receiver  string
	amount    float64
	signature string
}

type UTXO struct {
	TxID   string
	Index  int
	Amount float64
	Owner  string
}
