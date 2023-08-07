
/*
Encoding and decoding json using the encoding/json package
*/

package main

import (
	"encoding/json"
	"net/http"
	"fmt"
)

type User struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	Age int `json:"age"`
}

func main() {
	http.HandleFunc("/decode", func(w http.ResponseWriter, r *http.Request) {
		var user User
		json.NewDecoder(r.Body).Decode(&user)

		fmt.Fprintf(w, "%s %s is %d years old! \n", user.Firstname, user.Lastname, user.Age)
	})

	http.HandleFunc("/encode", func(w http.ResponseWriter, r *http.Request) {
		bond := User {
			Firstname: "James",
			Lastname: "Bond",
			Age: 50,
		}

		json.NewEncoder(w).Encode(bond)
	})

	http.ListenAndServe(":8080", nil)
}
