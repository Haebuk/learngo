package main

import (
	"github.com/learngo/nomadcoin/blockchain"
	"github.com/learngo/nomadcoin/cli"
)



func main() {
	blockchain.Blockchain()
	cli.Start()
} 