package blockchain

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

// BlockChain stores a slice of Block references
type BlockChain []*Block

func CreateBlockChain() *BlockChain {
	genesisBlock := NewBlock([]*Transaction{NewTransaction("", "", 0, "")}, "")
	genesisBlock.HashBlock(0)
	bc := &BlockChain{genesisBlock}
	UpdateBlockChain(bc)
	return bc
}

// ReadBlockChain parses database.json into struct format and returns
// the resultant BlockChain
func ReadBlockChain() (*BlockChain, error) {
	filename := "database.json"

	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Printf("failed to open json file: %s, error: %v", filename, err)
		return nil, err
	}
	defer jsonFile.Close()

	jsonData, err := io.ReadAll(jsonFile)
	if err != nil {
		fmt.Printf("failed to read json file, error: %v", err)
		return nil, err
	}

	var b BlockChain
	err = json.Unmarshal([]byte(jsonData), &b)
	if err != nil {
		return nil, err
	}

	return &b, nil
}

// UpdateBlockChain writes the current BlockChain to database.json
// TODO: Perform this write without rewriting the entire file
func UpdateBlockChain(b *BlockChain) {
	bm, err := json.Marshal(b)
	if err != nil {
		fmt.Println(err)
	}

	err = os.WriteFile("database.json", bm, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func (b *BlockChain) AddBlock(payload []*Transaction) {
	// Can verify if block hash passes proof of work here
	prevBlock := b.GetLeadingBlock()
	blockToAdd := NewBlock(payload, prevBlock.Hash)
	blockToAdd.HashBlock(0)
	*b = append(*b, blockToAdd)
	UpdateBlockChain(b)
}

//func (b *BlockChain) RemoveBlock() {
//	b = b[:len(b)-2]
//	UpdateBlockChain(b)
//}

// GetLeadingBlock returns the most recent block of a BlockChain
func (b *BlockChain) GetLeadingBlock() *Block {
	return (*b)[len(*b)-1]
}

// Print prints an easily readable format of the blockchain to the terminal
func (b *BlockChain) Print() {
	for _, block := range *b {
		fmt.Printf("Prev Hash - %x\n", []byte(block.Hash))
		fmt.Printf("%+v\n", block.Timestamp)
		fmt.Printf("- Payload - \n")
		for _, p := range block.Payload {
			fmt.Printf("%+v\n", p)
		}
		fmt.Printf("Cur Hash - %x\n", []byte(block.Hash))
		fmt.Printf("\n")
	}
}
