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

type MainData struct {
	PageTitle, Info string
}

func main() {
	// Parse a template
	tmpl := template.Must(template.ParseFiles("layout.html"))
	mtpl := template.Must(template.ParseFiles("index.html"))

	r := mux.NewRouter() // create a router

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := MainData {
			PageTitle: "This is the way",
			Info: "The Mandalorian",
		}

		mtpl.Execute(w, data)
	})

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

	http.ListenAndServe(":8080", r)
}
