package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"html/template" // package provides rich templating language for html templates
)

type Todo struct {
	Title string
	Done bool
}

type TodoPageData struct {
	PageTitle string
	Todos []Todo
}

func Templates() {
	// Parse a template
	tmpl := template.Must(template.ParseFiles("layout.html"))

	r := mux.NewRouter() // create a router

	// execute a template in a request handle
	r.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		// define data (A todo list & page title to be passed into a template)
		data := TodoPageData{
			PageTitle: "My TODO list",
			Todos: []Todo{
					{Title: "Task 1", Done: false},
					{Title: "Task 2", Done: true},
					{Title: "Task 3", Done: true},
			},
	}
		tmpl.Execute(w, data)
	})

	http.ListenAndServe(":80", r)
}
