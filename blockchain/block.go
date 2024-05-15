package blockchain

import (
	"math/rand"
	"time"
)
 
type Block struct {
	Hash     string
	Data     string
	PrevHash string
	Nonce 	 int
}

func CreateBlock(data string, prevHash string) *Block {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator
	initialNonce := rand.Intn(10000)
	block := &Block{"", data, prevHash, initialNonce}
	newPow := NewProofOfWork(block)
	
	nonce, hash := newPow.MineBlock()
	block.Hash = string(hash[:])
	block.Nonce = nonce
	return block
}

func Genesis() *Block {
	return CreateBlock("Genesis", "")
}
