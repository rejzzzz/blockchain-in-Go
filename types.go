package main

import(
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

type wallet struct{
	privateKey *ecdsa.PrivateKey
	publicKey []byte
}

type transaction struct{
	sender string
	receiver string
	amount float64
}