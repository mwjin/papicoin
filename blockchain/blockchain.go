package blockchain

import (
	"crypto/sha256"
	"fmt"
	"sync"
)

type block struct {
	data     string
	hash     string
	prevHash string
}

func (b *block) calculateHash() {
	b.hash = fmt.Sprintf("%x", sha256.Sum256([]byte(b.data+b.prevHash)))
}

type blockchain struct {
	blocks []*block
}

func (b *blockchain) AddBlock(data string) {
	b.blocks = append(b.blocks, createBlock(data))
}

func (b *blockchain) ListBlocks() {
	for _, blk := range b.blocks {
		fmt.Printf("Data: %s\n", blk.data)
		fmt.Printf("Hash: %s\n", blk.hash)
		fmt.Printf("Prev. Hash: %s\n", blk.prevHash)
	}
}

var b *blockchain
var once sync.Once

func getLastHash() string {
	totalBlocks := len(b.blocks)

	if totalBlocks == 0 {
		return ""
	}
	return b.blocks[totalBlocks-1].hash
}

func createBlock(data string) *block {
	newBlock := &block{data, "", getLastHash()}
	newBlock.calculateHash()
	return newBlock
}

func GetBlockchain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{}
			b.AddBlock("Genesis Block")
		})
	}
	return b
}
