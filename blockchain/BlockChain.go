package blockchain

import "fmt"

type BlockChain struct {
	blocks []*Block
}

func NewBlockChain() *BlockChain {
	genesisBlock := NewBlock([]byte("Genesis"), []byte{})
	genesisBlock.HashBlock(0)
	return &BlockChain{
		blocks: []*Block{genesisBlock},
	}
}

func (b *BlockChain) AddBlock(payload []byte) *BlockChain {
	// Can verify if block hash passes proof of work here
	prevBlock := b.GetLeadingBlock()
	blockToAdd := NewBlock(payload, prevBlock.hash)
	blockToAdd.HashBlock(0)
	b.blocks = append(b.blocks, blockToAdd)
	return b
}

func (b *BlockChain) GetLeadingBlock() *Block {
	return b.blocks[len(b.blocks)-1]
}

func (b *BlockChain) Print() {
	for _, block := range b.blocks {
		fmt.Printf("Prev Hash - %x\n", block.prevHash)
		fmt.Printf("Payload - %s\n", block.payload)
		fmt.Printf("Cur Hash - %x\n", block.hash)
		fmt.Printf("\n")
	}
}
