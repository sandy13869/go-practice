package main

import (
	"fmt"
	"log"

	"go-blockchain/wallet"
)

func init() {
	log.SetPrefix("Sam_Blockchain: ")
}

func main() {
	log.Println("Blockchain Started...")
	w := wallet.NewWallet()
	fmt.Println(w.PrivateKeyStr())
	fmt.Println(w.PublicKeyStr())
}
