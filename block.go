package main

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
)

type block struct {
	index        int
	timestamp    string
	data         string
	Nonce        int
	previousHash string
	hash         string
}

func (B *block) calculateHash() string {

	record := strconv.Itoa(B.index) + B.timestamp + B.data + B.previousHash + strconv.Itoa(B.Nonce)
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func (B *block) mineBlock(difficulty int) {
	prefix := ""
	for i := 0; i < difficulty; i++ {
		prefix += "0"
	}

	for {
		B.hash = B.calculateHash()
		if B.hash[:difficulty] == prefix {
			break
		}
		B.Nonce++
	}
}
