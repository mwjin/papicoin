package main

import (
	"github.com/mwjjeong/papicoin/explorer"
	"github.com/mwjjeong/papicoin/rest"
)

func main() {
	go rest.Start(4000)
	explorer.Start(3000)
}
