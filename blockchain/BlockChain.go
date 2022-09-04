package blockchain

import "fmt"

type BlockChain []*Block

func NewBlockChain() *BlockChain {
	genesisBlock := NewBlock(TransactionList{NewTransaction("", "", 0, "")}, []byte{})
	genesisBlock.HashBlock(0)
	return &BlockChain{genesisBlock}
}

func (b *BlockChain) AddBlock(payload TransactionList) {
	// Can verify if block hash passes proof of work here
	prevBlock := b.GetLeadingBlock()
	blockToAdd := NewBlock(payload, prevBlock.hash)
	blockToAdd.HashBlock(0)
	*b = append(*b, blockToAdd)
}

func (b *BlockChain) GetLeadingBlock() *Block {
	return (*b)[len(*b)-1]
}

func (b *BlockChain) Print() {
	for _, block := range *b {
		fmt.Printf("Prev Hash - %x\n", block.prevHash)
		fmt.Printf("%+v\n", block.timestamp)
		fmt.Printf("- Payload - \n")
		for _, p := range block.payload {
			fmt.Printf("%+v\n", p)
		}
		fmt.Printf("Cur Hash - %x\n", block.hash)
		fmt.Printf("\n")
	}
}
