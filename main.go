package main

import (
	"fmt"

	"github.com/Haebuk/learngo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "first word"}
	definition, err := dictionary.Search("Second")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}
}