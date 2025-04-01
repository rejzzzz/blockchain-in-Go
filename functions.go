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
	"math/big"
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

func createTransaction(sender *wallet, receiver []byte, amount float64) transaction{
	tx := transaction{
		sender : hex.EncodeToString(sender.publicKey),
		receiver : hex.EncodeToString(receiver),
		amount : amount,
	}

	// sign transaction
	data := fmt.Sprintf("%s%s%f", tx.sender, tx.receiver, tx.amount)
	hash := sha256.Sum256([]byte(data))
	r, s, _ := ecdsa.Sign(rand.Reader, sender.privateKey, hash[:])
	signature := append(r.Bytes(), s.Bytes()...)
	tx.signature = hex.EncodeToString((signature))

	return tx;
}

func verifyTransaction(tx *transaction) bool{
	data := fmt.Sprintf("%s%s%f", tx.sender, tx.receiver, tx.amount)
	hash := sha256.Sum256([]byte(data))

	sigBytes, _ := hex.DecodeString(tx.signature)
	r := big.Int{}
	s := big.Int{}
	r.SetBytes(sigBytes[:len(sigBytes)/2])
	s.SetBytes(sigBytes[len(sigBytes)/2:])

	publicKeyBytes, _ := hex.DecodeString(tx.sender)
	x := big.Int{}
	y := big.Int{}
	x.SetBytes(publicKeyBytes[:len(publicKeyBytes)/2])
	y.SetBytes(publicKeyBytes[len(publicKeyBytes)/2:])

	publicKey := ecdsa.PublicKey{Curve: elliptic.P256(), X: &x, Y: &y}
	return ecdsa.Verify(&publicKey, hash[:], &r, &s)
}

func updateBalance(block block){
	for _, tx := range block.transaction{
		if verifyTransaction(tx){
			balances[tx.sender] -= tx.amount  
			balances[tx.receiver] += tx.amount
		}
	}
}