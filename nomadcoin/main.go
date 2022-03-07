package main

import (
	"fmt"
	"os"
)

func usage() {
	fmt.Printf("Welcome to 노마드 코인\n\n")
	fmt.Println("Please use the following commands:")
	fmt.Println("explorer:	Start the HTML Explorer")
	fmt.Println("rest:		Start the REST API(recommended)")
	os.Exit(0)
}


func main() {
	if len(os.Args) < 2 {
		usage()
	}
	switch os.Args[1] {
	case "explorer":
		fmt.Println("Start Explorer")
	case "rest":
		fmt.Println("Start REST API")
	default:
		usage()
	}
 } 