package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/mwjjeong/papicoin/explorer"
	"github.com/mwjjeong/papicoin/rest"
)

func Start() {
	mode := flag.String("mode", "rest", "Choose between 'html', 'rest' and 'both'")
	port1 := flag.Int("port", 3000, "Set primary port of the server")
	port2 := flag.Int("port2", 4000, "Set secondary port of the server")

	if len(os.Args) < 2 {
		printUsage()
		os.Exit(0)
	}
	flag.Parse()
	run(*mode, *port1, *port2)
}

func printUsage() {
	fmt.Printf("Welcome to Papicoin!\n\n")
	fmt.Printf("Please use the following flags.\n\n")
	flag.PrintDefaults()
}

func run(mode string, port1 int, port2 int) {
	switch mode {
	case "explorer":
		fmt.Printf("Run Papicoin Explorer on %d\n\n", port1)
		explorer.Start(port1)
	case "rest":
		fmt.Printf("Run Papicoin REST API server on %d\n\n", port1)
		rest.Start(port1)
	case "both":
		fmt.Printf("Run Papicoin Explorer(%d) and REST API(%d)\n\n", port1, port2)
		go explorer.Start(port1)
		rest.Start(port2)
	default:
		printUsage()
	}
}
