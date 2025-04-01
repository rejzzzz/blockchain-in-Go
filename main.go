package main

import (
	"fmt"
)

var blockchain []block
var balances = make(map[string]float64)
var utxoSet = []UTXO{}


func main() {
	// first block
	firstBlock := createFirstBlock()
	blockchain = append(blockchain, firstBlock)

	fmt.Println("First Block Created: ")
	fmt.Printf("%+v\n\n", firstBlock)

	// add new blocks
	addBlock("Second Block creation")
	fmt.Println("Block 2 created, 1st addition:")
	fmt.Printf("%+v\n\n", blockchain[1])

	addBlock("third Block creation")
	fmt.Println("Block 3 created, 2nd addition:")
	fmt.Printf("%+v\n\n", blockchain[2])

	fmt.Println("Is blockchain valid? ", isBlockchainValid())


	

}
