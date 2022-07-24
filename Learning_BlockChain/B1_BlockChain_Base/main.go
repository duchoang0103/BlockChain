package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

type BlockChain struct {
	blocks []*Block
}

type Block struct {
	Timestamp int64
	Data      []byte
	Hash      []byte
	PrevHash  []byte
}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	info := bytes.Join([][]byte{b.Data, b.PrevHash, timestamp}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func NewBlock(data string, prevHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), []byte{}, prevHash}
	block.SetHash()
	return block
}

func (bc *BlockChain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	new := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, new)
}

func GenesisBlock() *Block {
	return NewBlock("Genesis", []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{GenesisBlock()}}
}

func main() {

	chain := InitBlockChain()
	chain.AddBlock("Hoang Thuan")
	chain.AddBlock("Hoang Luan")
	chain.AddBlock("Hoang Minh")
	for _, block := range chain.blocks {
		t := time.Unix(time.Now().Unix(), block.Timestamp)
		fmt.Printf("Time: ")
		fmt.Println(t.Format("02/01/2006, 15:04:05"))
		fmt.Printf("PrevHash: %x\n", block.PrevHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n\n", block.Hash)
	}
}
