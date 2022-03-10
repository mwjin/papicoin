package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/mwjjeong/papicoin/blockchain"
	"github.com/mwjjeong/papicoin/utils"
)

const (
	protocol string = "http://"
	domain   string = "localhost"
	port     string = ":4000"
)

type URL string

func (u URL) MarshalText() ([]byte, error) {
	return []byte(fmt.Sprintf("%s%s%s%s", protocol, domain, port, u)), nil
}

type URLDescription struct {
	URL         URL    `json:"url"`
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"`
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []URLDescription{
		{
			URL:         URL("/"),
			Method:      "GET",
			Description: "See Documentation",
		},
		{
			URL:         URL("/blocks"),
			Method:      "POST",
			Description: "Add a Block",
			Payload:     "data:string",
		},
	}
	rw.Header().Add("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(data)
}

type BlocksPostReqBody struct {
	Message string
}

func blocks(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		rw.Header().Add("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(blockchain.GetBlockchain().GetAllBlocks())
	case "POST":
		var reqBody BlocksPostReqBody
		utils.HandleErr(json.NewDecoder(r.Body).Decode(&reqBody))
		blockchain.GetBlockchain().AddBlock(reqBody.Message)
		rw.WriteHeader(http.StatusCreated)
	}
}
func main() {
	http.HandleFunc("/", documentation)
	http.HandleFunc("/blocks", blocks)
	addr := fmt.Sprintf("%s%s%s", protocol, domain, port)
	fmt.Printf("Listening on %s\n", addr)
	log.Fatal(http.ListenAndServe(port, nil))
}
