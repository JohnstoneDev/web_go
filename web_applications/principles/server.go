package main

import (
	"fmt"
	"strings"
	"net/http"
	"log"
)

/*
	Implementation of a basic server in Go
*/

// Simple requesthandler, prints some info on the path & a Hello string
func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()  // parse arguments, you have to call this by yourself
	fmt.Println(r.Form)  // print form information in server side

	fmt.Println("Path:", r.URL.Path)
	fmt.Println("Scheme:", r.URL.Scheme)

	fmt.Println("Form Value; url_long", r.Form["url_long"])
	for k, v := range r.Form {
			fmt.Println("key:", k)
			fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello There!") // send data to client side
}


// Entry Point, runs the server & listens for connections
func main() {
	http.HandleFunc("/", sayhelloName) // set router
	err := http.ListenAndServe(":9090", nil) // set listen port
	if err != nil {
			log.Fatal("ListenAndServe: ", err)
	}
}