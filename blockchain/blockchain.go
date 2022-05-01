package blockchain

import (
	"crypto/sha256"
	"fmt"
	"sync"
)

type Block struct {
	Data     string `json:"data"`
	Hash     string `json:"hash"`
	PrevHash string `json:"prevHash,omitempty"`
	Height   int    `json:"height"`
}

func (b *Block) calculateHash() {
	b.Hash = fmt.Sprintf("%x", sha256.Sum256([]byte(b.Data+b.PrevHash)))
}

type blockchain struct {
	blocks []*Block
}

func (b *blockchain) AddBlock(data string) {
	b.blocks = append(b.blocks, createBlock(data))
}

func (b *blockchain) PrintBlocks() {
	for _, blk := range b.blocks {
		fmt.Printf("Data: %s\n", blk.Data)
		fmt.Printf("Hash: %s\n", blk.Hash)
		fmt.Printf("Prev. Hash: %s\n", blk.PrevHash)
	}
}

func (b *blockchain) GetAllBlocks() []*Block {
	return b.blocks
}

func (b *blockchain) GetBlock(height int) *Block {
	return b.blocks[height-1]
}

var b *blockchain
var once sync.Once

func getLastHash() string {
	totalBlocks := len(b.blocks)

	if totalBlocks == 0 {
		return ""
	}
	return b.blocks[totalBlocks-1].Hash
}

func getCurrentHeight() int {
	totalBlocks := len(b.blocks)

	if totalBlocks == 0 {
		return 0
	}
	return b.blocks[totalBlocks-1].Height
}

func createBlock(data string) *Block {
	newBlock := &Block{data, "", getLastHash(), getCurrentHeight() + 1}
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
