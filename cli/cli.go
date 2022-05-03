package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/mwjjeong/papicoin/explorer"
	"github.com/mwjjeong/papicoin/rest"
)

func Start() {
	mode := flag.String("mode", "rest", "Choose between 'html' and 'rest'")
	port := flag.Int("port", 4000, "Set port of the server")

	if len(os.Args) < 2 {
		printUsage()
		os.Exit(0)
	}
	flag.Parse()
	run(*mode, *port)
}

func printUsage() {
	fmt.Printf("Welcome to Papicoin!\n\n")
	fmt.Printf("Please use the following flags.\n\n")
	flag.PrintDefaults()
}

func run(mode string, port int) {
	switch mode {
	case "explorer":
		fmt.Printf("Run Papicoin Explorer\n\n")
		explorer.Start(port)
	case "rest":
		fmt.Printf("Run Papicoin REST API server\n\n")
		rest.Start(port)
	default:
		printUsage()
	}
}
