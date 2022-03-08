package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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
			URL:         "/",
			Method:      "GET",
			Description: "See Documentation",
		},
	}
	rw.Header().Add("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(data)
}

func main() {
	http.HandleFunc("/", documentation)
	addr := fmt.Sprintf("http://localhost%s", port)
	fmt.Printf("Listening on %s\n", addr)
	log.Fatal(http.ListenAndServe(port, nil))
}
