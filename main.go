package main

import (
	"fmt"
	"os"

	"github.com/mwjjeong/papicoin/explorer"
	"github.com/mwjjeong/papicoin/rest"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(0)
	}
	run()
}

func printUsage() {
	fmt.Printf("Welcome to Papicoin!\n\n")
	fmt.Printf("Please use the following commands.\n\n")
	fmt.Println("explorer:	Start the HTML Explorer")
	fmt.Println("rest:		Start the REST API (recommended)")
}

func run() {
	switch os.Args[1] {
	case "explorer":
		fmt.Println("Start Explorer")
		explorer.Start(3000)
	case "rest":
		fmt.Println("Start REST API")
		rest.Start(4000)
	default:
		printUsage()
	}
}
