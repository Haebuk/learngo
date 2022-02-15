package main

import (
	"fmt"
	"log"

	"github.com/Haebuk/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("rjs")
	account.Deposit(10)
	fmt.Println(account.Balance())
	err := account.WithDraw(5)
	if err != nil{
		log.Fatalln(err)
	}
	fmt.Println(account.Balance())
	fmt.Println(account.Owner())
	fmt.Println(account)
	account.ChangeOwner("newrjs")
	fmt.Println(account)
}