package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/mwjjeong/papicoin/utils"
)

const port string = ":4000"

type URLDescription struct {
	URL         string
	Method      string
	Description string
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []URLDescription{
		{
			"/",
			"GET",
			"See Documentation",
		},
	}
	b, err := json.Marshal(data)
	utils.HandleErr(err)
	fmt.Fprintf(rw, "%s", b)
}

func main() {
	http.HandleFunc("/", documentation)
	addr := fmt.Sprintf("http://localhost%s", port)
	fmt.Printf("Listening on %s\n", addr)
	log.Fatal(http.ListenAndServe(port, nil))
}
