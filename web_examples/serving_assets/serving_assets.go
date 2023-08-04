package servingassets

import (
	"net/http"
)

/*
	to run :
		Change to main package
		go run serving_assets.go
		curl -s http://localhost:8080/static/css/styles.css
*/

func ServingAssets() {
	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/static", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8080", nil)
}
