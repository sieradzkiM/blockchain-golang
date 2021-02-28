package blockchain

import (
	"bytes"
	"crypto/sha256"
)

type Block struct {
	/*
		:Data - Keeps the record of all the completed transaction in the block.
		:Hash - Hash representation of the current block
		:PrevHash - Hash representation of the previous block
	*/
	Data     []byte
	Hash     []byte
	PrevHash []byte
}

type BlockChain struct {
	/*
		:Blocks - contain array of poiters to Blocks
	*/
	Blocks []*Block
}

func (b *Block) calculateHash() {
	record := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(record)
	b.Hash = hash[:]
}

func createBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte(data), []byte{}, prevHash}
	block.calculateHash()
	return block

}

func (blockchain *BlockChain) AddBlock(data string) {
	prevBlock := blockchain.Blocks[len(blockchain.Blocks)-1]
	newBlock := createBlock(data, prevBlock.Hash)
	blockchain.Blocks = append(blockchain.Blocks, newBlock)
}

func createGenesisBlock() *Block {
	return createBlock("", []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{createGenesisBlock()}}
}
