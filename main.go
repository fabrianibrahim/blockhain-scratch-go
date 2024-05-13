package main

import (
	"fmt"
	"blockChain/blockchain"
)

func main() {
	fmt.Println("Blockchain from Scratch")
	chain := blockchain.InitBlockChain()
	chain.AddBlock("Block 1")
	chain.AddBlock("Block 2")
	chain.AddBlock("Block 3")

	for _, block := range chain.Blocks {
		fmt.Printf("Previous hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash of block: %x\n", block.Hash)
		fmt.Println()

	}
}
