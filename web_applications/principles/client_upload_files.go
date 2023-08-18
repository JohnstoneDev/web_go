
// impersonating a form to upload a file
package main

import (
	"io"
	"os"
	"fmt"
	"bytes"
	"net/http"
	"mime/multipart"
)

// function that emulates posting a file by a client
func postFile(filename string, targetUrl string) error {
	bodyBuff := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuff)

	fileWriter, err := bodyWriter.CreateFormFile("uploadfile", filename)
	if err != nil {
		fmt.Println("Error Writing to buffer")
		return err
	}

	// open file handle
	fileHandle, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file")
		return err
	}

	defer fileHandle.Close()

	// iocopy
	_, er := io.Copy(fileWriter, fileHandle)
	if er != nil {
		fmt.Println("Error occurs in io.Copy()")
		return er
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	resp, err := http.Post(targetUrl, contentType, bodyBuff)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	resp_body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(resp.Status)
	fmt.Println(resp.StatusCode)

	fmt.Println(string(resp_body))
	return nil
}

func main() {
	targetUrl := "/upload"
	filename := "./uploaded_sorting_algo.png"
	postFile(filename, targetUrl)
}