package main

import (
	"simple_blockchain/blockchain"
)

func main() {
	bc := blockchain.NewBlockchain() // Create a new blockchain
	defer bc.Db.Close()

	cli := blockchain.CLI{Bc: bc} // Create a new CLI
	cli.Run()
}
