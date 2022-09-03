package main

import "fmt"

func testFunction(c chan string) {
	c <- "Hello Channel"
	c <- "String 2"
	c <- "String 3"
	c <- "String 4"
	c <- "String 5"
	c <- "String 6"
	c <- "String 7"
	c <- "String 8"
	c <- "String 9"
}

func main() {
	c := make(chan string, 2)

	go testFunction(c)

	for {
		fmt.Println(<-c)
		if len(c) == 0 {
			break
		}
	}
}

//chain := blockchain.NewBlockChain()
//chain.AddBlock([]byte("Hello"))
//chain.AddBlock([]byte("Hello1"))
//chain.AddBlock([]byte("Hello2"))
