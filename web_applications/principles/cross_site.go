package main

/*
	NOTES :
	To prevent XSS, yous hould combine :
		Validation of all data from users
		Crefully handle all data that will be sent ot clients in order to
		prevent any injected scripts from running on browsers

*/

/*
	func HTMLEscape(w io.Writer, b []byte) escapes b to w.
	func HTMLEscapeString(s string) string returns a string after escaping from s.
	func HTMLEscaper(args ...interface{}) string returns a string after escaping from multiple arguments.
*/

import (
	"fmt"
	"log"
	"net/http"
	"html/template"
	// "text/template" // => if you want to respond with text
)

// router for log in route
func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method: ", r.Method) // get request method

	if r.Method == "GET" {
		t := template.Must(template.ParseFiles("login.html")) // parse form & send
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		// logic to log in
    fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username"))) // print at server side
    fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))
    template.HTMLEscape(w, []byte(r.Form.Get("username"))) // responded to clients

		// t, err := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`) // use text template
		// err = t.ExecuteTemplate(out, "T", "<script>alert('you have been pwned')</script>")
	}
}

func main() {
	http.HandleFunc("/login", login)

	err := http.ListenAndServe(":9090", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
