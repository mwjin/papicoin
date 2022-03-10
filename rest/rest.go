package rest

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

type url string

func (u url) MarshalText() ([]byte, error) {
	return []byte(fmt.Sprintf("%s%s%s%s", protocol, domain, port, u)), nil
}

type urlDescription struct {
	URL         url    `json:"url"`
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"`
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []urlDescription{
		{
			URL:         url("/"),
			Method:      "GET",
			Description: "See Documentation",
		},
		{
			URL:         url("/blocks"),
			Method:      "GET",
			Description: "See the List of Blocks",
		},
		{
			URL:         url("/blocks"),
			Method:      "POST",
			Description: "Add a Block",
			Payload:     "data:string",
		},
	}
	rw.Header().Add("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(data)
}

type blocksPostReqBody struct {
	Message string
}

func blocks(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		rw.Header().Add("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(blockchain.GetBlockchain().GetAllBlocks())
	case "POST":
		var reqBody blocksPostReqBody
		utils.HandleErr(json.NewDecoder(r.Body).Decode(&reqBody))
		blockchain.GetBlockchain().AddBlock(reqBody.Message)
		rw.WriteHeader(http.StatusCreated)
	}
}
func Start() {
	http.HandleFunc("/", documentation)
	http.HandleFunc("/blocks", blocks)
	addr := fmt.Sprintf("%s%s%s", protocol, domain, port)
	fmt.Printf("Listening on %s\n", addr)
	log.Fatal(http.ListenAndServe(port, nil))
}
