package main

import (
	"github.com/learngo/nomadcoin/explorer"
	"github.com/learngo/nomadcoin/rest"
)


func main() {
	go explorer.Start(3000)
	rest.Start(4000)
 }