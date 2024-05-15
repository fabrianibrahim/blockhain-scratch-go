package main

import (
	"fmt"
	"strconv"
	"blockChain/blockchain"
)


func main() {
    chain := blockchain.InitBlockChain()
    aliceWallet, err := blockchain.NewWallet()
    
    if err != nil {
        fmt.Println("Error creating Alice's wallet:", err)
        return
    }
    fmt.Println("Alice's wallet created successfully")

    bobWallet, err := blockchain.NewWallet()    
    if err != nil {
        fmt.Println("Error creating Bob's wallet:", err)
        return
    }
    fmt.Println("Bob's wallet created successfully")

    tx := &blockchain.Transaction{
        Sender:   aliceWallet.PublicKey.N.String(),
        Receiver: bobWallet.PublicKey.N.String(),
        Amount:   5.0,
    }
    fmt.Println("Alice to Bob Transaction created successfully")

    signature, err := aliceWallet.SignTransaction(tx)
    if err != nil {
        fmt.Println("Error signing the transaction:", err)
        return
    }

    // Verify the transaction using Alice’s wallet, public key, and the signature.
    err = blockchain.VerifyTransaction(tx, aliceWallet.PublicKey, signature)
    
    if err != nil {
        fmt.Println("Transaction verification failed:", err)
        return
    }

    fmt.Println("Transaction Verified Successfully")

    // Add the verified transaction to the blockchain.
    chain.AddBlock("Block 1", "Alice", []*blockchain.Transaction{tx})
    fmt.Println()

    // Print the blockchain.
    for _, block := range chain.Blocks {
        fmt.Printf("Previous hash: %x\n", block.PrevHash)
        fmt.Printf("Data: %s\n", block.Data)
        fmt.Printf("Hash: %x\n", block.Hash)

        pow := blockchain.NewProofOfWork(block)
        fmt.Printf("IsValidPoW: %s\n", strconv.FormatBool(pow.Validate()))
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