package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/mwjjeong/papicoin/blockchain"
)

const (
	port        string = ":4000"
	templateDir string = "templates/"
)

var templates *template.Template

type PageData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

func home(rw http.ResponseWriter, r *http.Request) {
	data := PageData{"Home", blockchain.GetBlockchain().GetAllBlocks()}
	templates.ExecuteTemplate(rw, "home", data)
}

func add(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		templates.ExecuteTemplate(rw, "add", nil)
	case "POST":
		r.ParseForm()
		data := r.Form.Get("blockData")
		blockchain.GetBlockchain().AddBlock(data)
		fmt.Printf("Add the data '%s' to the chain.\n", data)
		http.Redirect(rw, r, "/", http.StatusPermanentRedirect)
	default:
		fmt.Println("Hello")
	}
}

func main() {
	templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.gohtml"))
	http.HandleFunc("/", home)
	http.HandleFunc("/add", add)
	fmt.Printf("Listerning on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
