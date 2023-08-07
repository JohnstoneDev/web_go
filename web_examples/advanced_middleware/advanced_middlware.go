package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

/*
	defining a type Middleware that makes it easier to chain multiple
	middleware together
*/


// define the middleware type
type MiddleWare func (w http.HandlerFunc) http.HandlerFunc

// Logging logs all requests with its path and the time it took to process
func Logging() MiddleWare {
	// create a new Middleware
	middleware := func (f http.HandlerFunc) http.HandlerFunc {
		// Define the http.HandlerFunc
		handler := func (w http.ResponseWriter, r *http.Request) {
			// ... do middleware stuff
			start := time.Now()

			defer func() { log.Println(r.URL.Path, time.Since(start))} ()

			// call the next middleware / handler
			f(w, r)
		}
		return handler
	}

	return middleware
}

// Method verifies that a url has been requested with a specific method, else returns error code 400
func Method(m string) MiddleWare {
	// create new middleware
	middleware := func (f http.HandlerFunc) http.HandlerFunc {
		// create a handler
		handler := func (w http.ResponseWriter, r *http.Request) {
			// do middleware stuff (verify that the request method id same as what is expected)
			if r.Method != m {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}
			// call the next middleware / handler
			f(w, r)
		}
		return handler
	}
	return middleware
}

// Chain applies middleware to a http.HandlerFunc

func Chain(f http.HandlerFunc, middlewares ...MiddleWare) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}

	return f
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello there \n")
}

func main() {
	http.HandleFunc("/", Chain(Hello, Method("GET"), Logging()))
	http.ListenAndServe(":8080", nil)
}