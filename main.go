package main

import (
	"fmt"
	"strconv"
	"blockChain/blockchain"
)

func main() {
	fmt.Println("Blockchain from Scratch")
	chain := blockchain.InitBlockChain()
	chain.AddBlock("Block 1", "Alice", []*blockchain.Transaction{
        {Sender: "Alice", Receiver: "Bob", Amount: 1.5},
        {Sender: "Alice", Receiver: "Charlie", Amount: 19.5},
    })

    chain.AddBlock("Block 2", "Bob", []*blockchain.Transaction{
        {Sender: "Bob", Receiver: "Charlie", Amount: 2.3},
    })

	for _, block := range chain.Blocks {
		fmt.Printf("Previous hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash of block: %x\n", block.Hash)

		pow := blockchain.NewProofOfWork(block)
		fmt.Printf("IsValidPow: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

		fmt.Println("Transactions:")

        for _, tx := range block.Transactions {
            fmt.Printf("Sender: %s\n", tx.Sender)
            fmt.Printf("Receiver: %s\n", tx.Receiver)
            fmt.Printf("Amount: %f\n", tx.Amount)
            fmt.Printf("Coinbase: %t\n", tx.Coinbase)
            fmt.Println()
        }
        fmt.Println()
	}
}
