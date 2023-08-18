/*
To prevent duplicate submissions, you can add a hidden field with
a unique token, and to always check this token before processing the
incoming data.

You can use an MD5 hash (time stamp) to generate the token, and add it to both
a hidden field on the client side form & a session cookie on the server.
You can then use this token to check whether or not this form was submitted
*/

package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)


func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method:", r.Method)

	if r.Method == "GET" {
		crutime := time.Now().Unix()

		hsh := md5.New()
		io.WriteString(hsh, strconv.FormatInt(crutime, 10))

		token := fmt.Sprintf("%x", hsh.Sum(nil))

		tmpl, _ := template.ParseFiles("duplicate_submissions.html")
		tmpl.Execute(w, token)
	} else {
		// log in request
		r.ParseForm()
		token := r.Form.Get("token")
		if token != "" {
				// check token validity
		} else {
				// give error if no token
		}
		fmt.Println("username length:", len(r.Form["username"][0]))
		fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username"))) // print in server side
		fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))
		template.HTMLEscape(w, []byte(r.Form.Get("username"))) // respond to client
	}
}

func main() {
	http.HandleFunc("/login", login)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
