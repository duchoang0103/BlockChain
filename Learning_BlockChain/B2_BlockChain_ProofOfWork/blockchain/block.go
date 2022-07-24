package blockchain

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type BlockChain struct {
	//blocks []*Block
	Blocks []*Block
}

type Block struct {
	Timestamp int64
	Data      []byte
	Hash      []byte
	PrevHash  []byte
	Nonce     int
}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	info := bytes.Join([][]byte{b.Data, b.PrevHash, timestamp}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func NewBlock(data string, prevHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), []byte{}, prevHash, 0}
	//block.SetHash()
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}

func (bc *BlockChain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	new := NewBlock(data, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, new)
}

func GenesisBlock() *Block {
	return NewBlock("Genesis", []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{GenesisBlock()}}
}
