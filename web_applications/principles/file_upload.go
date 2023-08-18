/*
	If your website/application would require users to upload files, you have
	to add the property **enctype** to the form that you want to use to upload

	There are three possible values for this property :
	application/x-www-form-urlencoded   Transcode all characters before uploading (default).
	multipart/form-data   No transcoding. You must use this value when your form has file upload controls.
	text/plain    Convert spaces to "+", but no transcoding for special characters.
*/

package main

import (
	"io"
	"os"
	"fmt"
	"time"
	"strconv"
	"net/http"
	"crypto/md5"
	"html/template"
)

// Request handler for file uploads : Prints file info to the server
func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method: ", r.Method)

	if r.Method == "GET" {
		// Generate token to prevent duplicate posting
		crutime := time.Now().Unix()
		h := md5.New()

		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("file_upload.html")
		t.Execute(w, token)
	} else {

		// file is  saved in the server memory with maxMemory size
		// if the file is larger than maxMemory, the rest of the data will be
		// saved in a system temp file,

		r.ParseMultipartForm(32 << 29) // takes the maxMemory argument

		// use r.FormFile to get the file handle & use io.Copy to save to the (your) filesystem
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}

		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)

		// opens the file & uploads it
		f, err := os.OpenFile("./uploaded_" + handler.Filename, os.O_WRONLY | os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}

		defer f.Close()
		io.Copy(f, file)
	}
}

func main() {
	http.HandleFunc("/upload", upload)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("ListenAndServe:", err)
	}
}