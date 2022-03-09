package main

import (
	"github.com/learngo/nomadcoin/blockchain"
	"github.com/learngo/nomadcoin/cli"
	"github.com/learngo/nomadcoin/db"
)



func main() {
	defer db.Close()
	blockchain.Blockchain()
	cli.Start()
} 