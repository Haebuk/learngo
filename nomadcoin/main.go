package main

import (
	"github.com/learngo/nomadcoin/cli"
	"github.com/learngo/nomadcoin/db"
)


 

func main() {
	// wallet.Wallet()
	cli.Start()
	defer db.Close()
} 