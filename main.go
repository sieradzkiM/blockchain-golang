package main

import (
	"fmt"
	"github.com/sieradzkim/blockchain-golang/blockchain"
	"strconv"
)

func main() {
	bchain := blockchain.InitBlockChain()
	bchain.AddBlock("1st Block")
	bchain.AddBlock("2nd Block")

	for i, block := range bchain.Blocks {
		fmt.Printf("Block index: %x\n", i)
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

		pow := blockchain.NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
