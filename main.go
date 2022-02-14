package main

import (
	"fmt"
	"strings"
)

func repeatMe(words ...string) {
	fmt.Println(words)
}

func lenAndUpper(name string) (length int, uppercase string){
	defer fmt.Println("I'm Done")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func main() {
	fmt.Println("Hello World")
	repeatMe("Hello", "World")
	lenth, uppercase := lenAndUpper("Hello World")
	fmt.Println(lenth, uppercase)
}
