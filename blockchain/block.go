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
	Transactions []*Transaction
}

func CreateBlock(data string, prevHash string, transactions []*Transaction) *Block {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator
	initialNonce := rand.Intn(10000)
	block := &Block{"", data, prevHash, initialNonce, transactions}
	newPow := NewProofOfWork(block)
	
	nonce, hash := newPow.MineBlock()
	block.Hash = string(hash[:])
	block.Nonce = nonce
	return block
}

func Genesis() *Block {
	coinbaseTransaction := &Transaction{
		Sender: "Coinbase",
		Receiver: "Genesis",
		Amount: 0.0,
		Coinbase: true,
	}
	return CreateBlock("Genesis", "", []*Transaction{coinbaseTransaction})
}
