package cli

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/learngo/nomadcoin/explorer"
	"github.com/learngo/nomadcoin/rest"
)

func usage() {
	fmt.Printf("Welcome to 노마드 코인\n\n")
	fmt.Println("Please use the following flags:")
	fmt.Println("-port:	Set the PORT of the server")
	fmt.Printf("-mode:	Choose between 'html' and 'rest'\n\n")
	runtime.Goexit()
}

func Start() {
	if len(os.Args) == 1 {
		usage()
	}

	port := flag.Int("port", 4000, "Set port of the server")
	mode := flag.String("mode", "rest", "Choose between 'html' and 'rest'")

	flag.Parse()

	switch *mode {
	case "rest":
		rest.Start(*port)
	case "html":
		explorer.Start(*port)
	default:
		usage()
	}
}