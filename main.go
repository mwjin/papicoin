package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const port string = ":4000"

type PageData struct {
	PageTitle string
}

func home(rw http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles("templates/home.html")
	page := PageData{"Home"}
	if err != nil {
		log.Fatal(err)
	}
	tmp.Execute(rw, page)
}

func main() {
	http.HandleFunc("/", home)
	fmt.Printf("Listerning on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
