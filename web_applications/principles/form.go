package main

import (
	"fmt"
	"html/template"
	"net/http"

	// "strings"
	"log"
)

// router for index route
func sayHello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	for k, v := range r.Form {
		fmt.Println("Key:", k)
		fmt.Println("Value:", v)
	}

	fmt.Fprintf(w, "Hello There!") // write data to response
}

// router for log in route
func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method: ", r.Method) // get request method

	if r.Method == "GET" {
		t := template.Must(template.ParseFiles("login.html")) // parse form & send
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		// logic to log in
		fmt.Println("username: ", r.Form["username"])
		fmt.Println("password: ", r.Form["password"])
	}
}

func main() {
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/login", login)

	err := http.ListenAndServe(":9090", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
