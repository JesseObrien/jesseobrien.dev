package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"

	"github.com/rakyll/statik/fs"

	_ "github.com/jesseobrien/jesseobrien.dev/statik"
)

var httpPort string
var httpAddress string

func init() {
	flag.StringVar(&httpPort, "port", "8080", "The port you wish to bind the server to.")
	flag.StringVar(&httpAddress, "ip", "127.0.0.1", "The IP address to bind to")

	flag.Parse()
}

func main() {

	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(statikFS)))

	r, err := statikFS.Open("/index.html")
	if err != nil {
		log.Fatal(err)
	}

	defer r.Close()

	indexTmpl, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	tmpl, err := template.New("index").Parse(string(indexTmpl))
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		tmpl.Execute(w, nil)
	})

	listenAddr := fmt.Sprintf("%s:%s", httpAddress, httpPort)
	log.Printf("Starting HTTP server on %s", listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}
