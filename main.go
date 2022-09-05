package main

import (
	"fmt"
	"mbc.mikesprogram.com/blockchain"
)

func main() {
	//chain := blockchain.CreateBlockChain()
	//chain.AddBlock([]*blockchain.Transaction{
	//	blockchain.NewTransaction("mike", "apple", 5, ""),
	//	blockchain.NewTransaction("apple", "kane", 5, ""),
	//	blockchain.NewTransaction("mike", "kane", 3, ""),
	//	blockchain.NewTransaction("kane", "mike", 1, ""),
	//})
	//
	//chain.AddBlock([]*blockchain.Transaction{
	//	blockchain.NewTransaction("mike", "kane", 1, ""),
	//	blockchain.NewTransaction("apple", "kane", 12, ""),
	//	blockchain.NewTransaction("mike", "kane", 3, ""),
	//	blockchain.NewTransaction("kane", "apple", 7, ""),
	//})
	//
	//chain.AddBlock([]*blockchain.Transaction{
	//	blockchain.NewTransaction("kane", "apple", 3, ""),
	//	blockchain.NewTransaction("mike", "apple", 15, ""),
	//	blockchain.NewTransaction("mike", "kane", 1, ""),
	//	blockchain.NewTransaction("apple", "mike", 2, ""),
	//})

	chain, err := blockchain.ReadBlockChain()
	if err != nil {
		fmt.Println(err)
		return
	}

	chain.Print()
}
