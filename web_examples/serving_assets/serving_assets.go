package servingassets

import (
	"net/http"
)

func ServingAssets() {
	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/static", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8080", nil)
}
