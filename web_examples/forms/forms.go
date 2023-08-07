package main


import (
	"html/template"
	"net/http"
	"github.com/gorilla/mux"
)

type ContactDetails struct {
	Email, Subject, Message string
}


func main() {
	router := mux.NewRouter() // register a router with mux

	tmpl := template.Must(template.ParseFiles("forms.html")) // parse a template

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		details := ContactDetails {
			Email : r.FormValue("email"),
			Subject: r.FormValue("subject"),
			Message: r.FormValue("message"),
		}

		// do something with details
		_ = details

		tmpl.Execute(w, struct{ Success bool}{true})
	})

	http.ListenAndServe(":8080", router)
}
