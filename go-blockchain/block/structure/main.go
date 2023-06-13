package main

import (
	"fmt"
	"log"
	"strings"
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

type Blockchain struct {
	transactionPool []string
	chain           []*Block
}

func NewBlockchain() *Blockchain {
	bc := new(Blockchain)
	bc.CreateBlock(0, "Init hash")
	return bc
}

func (bc *Blockchain) CreateBlock(nonce int, previousHash string) *Block {
	b := NewBlock(nonce, previousHash)
	bc.chain = append(bc.chain, b)
	return b
}

func (bc *Blockchain) Print() {
	for i, block := range bc.chain {
		fmt.Printf("%s Chain %d %s\n", strings.Repeat("=", 9), i, strings.Repeat("=", 9))
		block.Print()
	}
	fmt.Printf("%s \n", strings.Repeat("*", 25))
}

func init() {
	log.SetPrefix("Shandy_Blockchain_Network: ")
}

func main() {
	// log.Println("Application Started...")
	// fmt.Println("Welcome Sandy ...!")
	blockChain := NewBlockchain()
	// fmt.Println(blockChain)
	blockChain.Print()
	blockChain.CreateBlock(6, "hash 1")
	blockChain.Print()
	blockChain.CreateBlock(9, "hash 2")
	blockChain.Print()
}
