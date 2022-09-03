package main

import "mbc.mikesprogram.com/blockchain"

func main() {
	chain := blockchain.NewBlockChain()
	chain.AddBlock([]byte("Hello"))
	chain.AddBlock([]byte("Hello1"))
	chain.AddBlock([]byte("Hello2"))

	chain.Print()
}
