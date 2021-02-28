package blockchain

type Block struct {
	/*
		:Data - Keeps the record of all the completed transaction in the block.
		:Hash - Hash representation of the current block
		:PrevHash - Hash representation of the previous block
	*/
	Data     []byte
	Hash     []byte
	PrevHash []byte
	Nonce    int
}

type BlockChain struct {
	/*
		:Blocks - contain array of poiters to Blocks
	*/
	Blocks []*Block
}

func createBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte(data), []byte{}, prevHash, 0}
	//block.calculateHash()
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce

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
