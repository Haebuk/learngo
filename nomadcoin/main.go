package main

import (
	"github.com/learngo/nomadcoin/cli"
	"github.com/learngo/nomadcoin/db"
)

func main() {
	defer db.Close()
	cli.Start()
} 