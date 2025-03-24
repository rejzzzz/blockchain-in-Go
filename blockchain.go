package main

import (
	//"fmt"
	"fmt"
	"time"
)

var blockchain []block

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

func addBlock(Data string){
	previousBlock:= blockchain[len(blockchain)-1]
	newBlock:= block{
		index: previousBlock.index +1,
		timestamp: time.Now().String(),
		data: Data,
		previousHash: previousBlock.hash,
	}
	newBlock.mineBlock(5) // 5 is the difficulty
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

