package main

import (
	"crypto/sha256"
	"fmt"
)

type block struct {
	data     string
	hash     string
	prevHash string
}

func main() {
	genesisBlock := block{"This is a genesis block.", "", ""}
	hash := sha256.Sum256([]byte(genesisBlock.data + genesisBlock.prevHash))
	genesisBlock.hash = fmt.Sprintf("%x", hash)
	fmt.Println(genesisBlock.hash)
}