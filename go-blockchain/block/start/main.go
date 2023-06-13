package main

import (
	"fmt"
	"log"
	"time"
)

type Block struct {
	nonce        int
	previousHash string
	timestamp    int64
	transactions []string
}

func NewBlock(nonce int, previousHash string) *Block {
	b := new(Block)
	b.timestamp = time.Now().UnixNano()
	b.nonce = nonce
	b.previousHash = previousHash
	return b

	// return &Block{
	// 	timestamp: time.Now().UnixNano(),
	// }
}

func (b *Block) Print() {
	fmt.Printf("timestamp %d\n", b.timestamp)
	fmt.Printf("nonce %d\n", b.nonce)
	fmt.Printf("previous_hash %s\n", b.previousHash)
	fmt.Printf("transactions %s\n", b.transactions)
}

func init() {
	log.SetPrefix("Shandy_Blockchain_Network: ")
}

func main() {
	log.Println("Application Started...")
	fmt.Println("Welcome Sandy ...!")
	b := NewBlock(0, "init hash")
	b.Print()
}
