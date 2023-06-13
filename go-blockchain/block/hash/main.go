package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
)

type Block struct {
	nonce        int
	previousHash [32]byte
	timestamp    int64
	transactions []string
}

func NewBlock(nonce int, previousHash [32]byte) *Block {
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
	fmt.Printf("previous_hash %x\n", b.previousHash)
	fmt.Printf("transactions %s\n", b.transactions)
}

func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	fmt.Println(string(m))
	return sha256.Sum256([]byte(m))
}

func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Timestamp    int64    `json:"timestamp"`
		Nonce        int      `json:"nonce"`
		PreviousHash [32]byte `json:"previoushash"`
		Transactions []string `json:"transactions"`
	}{
		Timestamp:    b.timestamp,
		Nonce:        b.nonce,
		PreviousHash: b.previousHash,
		Transactions: b.transactions,
	})
}

type Blockchain struct {
	transactionPool []string
	chain           []*Block
}

func NewBlockchain() *Blockchain {
	b := &Block{}
	bc := new(Blockchain)
	bc.CreateBlock(0, b.Hash())
	return bc
}

func (bc *Blockchain) CreateBlock(nonce int, previousHash [32]byte) *Block {
	b := NewBlock(nonce, previousHash)
	bc.chain = append(bc.chain, b)
	return b
}

func (bc *Blockchain) LastBlock() *Block {
	return bc.chain[len(bc.chain)-1]
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
	// sum := sha256.Sum256([]byte("Hello Sandy!"))
	// block := &Block{nonce: 1}
	// fmt.Printf("%x \n", block.Hash())
	// fmt.Printf("%x \n", sum)
	blockChain := NewBlockchain()
	blockChain.Print()

	previousHash := blockChain.LastBlock().Hash()
	blockChain.CreateBlock(6, previousHash)
	blockChain.Print()

	previousHash = blockChain.LastBlock().Hash()
	blockChain.CreateBlock(9, previousHash)
	blockChain.Print()
}
