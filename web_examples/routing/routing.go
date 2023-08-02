package routing

import (
	"fmt"
	"net/http"


	"github.com/gorilla/mux"
)

/*
	1. First Install gorilla/ mux with go get -u github.com/gorilla/mux
	2. Create a new request router thar will receive all HTTP connections and pass
	it on to the request handlers you will register on it
	3. Extract segments from the URL / Request with gorilla / mux
	4. Perform some operations with data from the request
	5. Listen for connections on your router
*/

func Routing() {
	r := mux.NewRouter() // create a new request router

	// Create a request handler, match it to a dynamic route with placeholders
	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {

	// get the data from the segments of the url (will be saved in a map : vars)
	// mux.Vars() takes a parameter of the https.Request(here @r)
		vars := mux.Vars(r)

		title, page := vars["title"], vars["page"] // extract the variables from the map

		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	})

	// set the servers router (replace the nil with your request router)
	http.ListenAndServe(":80", r)
}