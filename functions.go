package main

import(
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"time"
	"fmt"
	"crypto/sha256"
	"encoding/hex"
	"strconv"
)

func (B *block) calculateHash() string {

	record := strconv.Itoa(B.index) + B.timestamp + B.data + B.previousHash + strconv.Itoa(B.Nonce)
	for _, tx := range B.transaction{
		record += tx.sender + tx.receiver + strconv.FormatFloat(tx.amount, 'f', -1, 64)
	}
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

// Genesis Block
func createFirstBlock() block{
	genesisBlock := block{
		index: 0,
		timestamp: time.Now().String(),
		data: "Genesis block",
		previousHash: "",
	}
	genesisBlock.mineBlock(5)

	return genesisBlock
}

// Blockchain   
var blockchain []block
func addBlock(Data string){
	previousBlock:= blockchain[len(blockchain)-1]
	newBlock:= block{
		index: previousBlock.index +1,
		timestamp: time.Now().String(),
		data: Data,
		previousHash: previousBlock.hash,
	}
	newBlock.mineBlock(5) // 5 is the difficulty or salt
	blockchain = append(blockchain, newBlock)
}

func isBlockchainValid() bool{
	for i:= 1; i<len(blockchain); i++{
		currentBlock:= blockchain[i]
		previousBlock:= blockchain[i-1]

		// check current block hash
		if currentBlock.hash != currentBlock.calculateHash(){
			fmt.Println("Invalid hash for current block", currentBlock.index)
			return false
		}
		
		// check previous block hash
		if currentBlock.previousHash != previousBlock.hash{
			fmt.Println("new block's previous hash value fdoes not match actual previous block's hash value", currentBlock.index)
			return false
		}
	}
	return true
}

func newWallet() *wallet{
	privateKey, _ := ecdsa.GenerateKey((elliptic.P256()), rand.Reader)
	publicKey := append(privateKey.PublicKey.X.Bytes(), privateKey.PublicKey.Y.Bytes()...)
	return &wallet{privateKey, publicKey}
}