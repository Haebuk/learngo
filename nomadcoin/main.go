package main

import (
	"flag"
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

	rest := flag.NewFlagSet("rest", flag.ExitOnError)

	portFlag := rest.Int("port", 4000, "Sets the port of the server")

	switch os.Args[1] {
	case "explorer":
		fmt.Println("Start Explorer")
	case "rest":
		fmt.Println("Start REST API")
		rest.Parse(os.Args[2:])
	default:
		usage()
	}
 } 