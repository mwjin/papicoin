package explorer

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/mwjjeong/papicoin/blockchain"
)

const (
	templateDir string = "explorer/templates/"
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

func Start(port int) {
	parseTemplate()
	registerHandlerFunc()
	fmt.Printf("Listerning on http://localhost:%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func parseTemplate() {
	templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.gohtml"))
}

func registerHandlerFunc() {
	http.HandleFunc("/", home)
	http.HandleFunc("/add", add)
}
