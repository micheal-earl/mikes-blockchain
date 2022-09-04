package main

import (
	"mbc.mikesprogram.com/blockchain"
)

func main() {
	chain := blockchain.NewBlockChain()
	chain.AddBlock(blockchain.TransactionList{
		blockchain.NewTransaction("mike", "jim", 5, ""),
		blockchain.NewTransaction("mike", "jim", 5, ""),
		blockchain.NewTransaction("mike", "jim", 5, ""),
		blockchain.NewTransaction("mike", "jim", 5, ""),
	})
	chain.AddBlock(blockchain.TransactionList{
		blockchain.NewTransaction("mike", "jim", 5, ""),
		blockchain.NewTransaction("mike", "jim", 5, ""),
		blockchain.NewTransaction("mike", "jim", 5, ""),
		blockchain.NewTransaction("mike", "jim", 5, ""),
	})
	chain.AddBlock(blockchain.TransactionList{
		blockchain.NewTransaction("mike", "jim", 5, ""),
		blockchain.NewTransaction("mike", "jim", 5, ""),
		blockchain.NewTransaction("mike", "jim", 5, ""),
		blockchain.NewTransaction("mike", "jim", 5, ""),
	})

	chain.Print()
}
