package hello

import (
	"fmt"
	"net/http"
)

/*
	a handler receives two parameters, the ResponseWRiter to write text / html responses
	a Request which contains all information about this HTTP request (URL ? header fields)
*/

func Hello() {
	// Reggister a request handler that receives all incoming HTTP Connections
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	// Listen to connections on the port
	http.ListenAndServe(":80", nil)
}
