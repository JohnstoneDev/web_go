package serve

import (
	"fmt"
	"net/http"
)

/*
	basics of creating a server in Go, should :
	1. Process Dynamic Requests
	2. Serve static assets
	3. Accept connections
*/


func Serve() {
	// register a new handler : takes the path to match & function ot execute (Dynamic Request)
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my website!")
	})

	// Serving static assets
	fs := http.FileServer(http.Dir("static/")) // use inbuilt FileServer & point to the url path(to serve files from)
	http.Handle("/static/", http.StripPrefix("/static/", fs)) // to serve file, strip url prefix(the name of the dir where tehe files are)

	// Accept connections
	http.ListenAndServe(":80", nil)
}