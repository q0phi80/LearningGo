package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

func postFileAndReturnResponse(filename string) string {
	// Create a buffer we can write the file to
	fileDataBuffer := bytes.Buffer{}
	multipartWritter := multipart.NewWriter(&fileDataBuffer)
	// Open the file from local computer to upload
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	// Once local file is opened, create a formFile. This wraps the file data in the right format to upload it to the server.
	formFile, err := multipartWritter.CreateFormFile("myFile", file.Name())
	if err != nil {
		log.Fatal(err)
	}
	// Copy the bytes from the local file into the form file, then close the form file writer so that it knows no more data will be added.
	_, err = io.Copy(formFile, file)
	if err != nil {
		log.Fatal(err)
	}
	multipartWritter.Close()

	req, err := http.NewRequest("POST", "http://localhost:8080", &fileDataBuffer)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", multipartWritter.FormDataContentType())
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}
func main() {
	data := postFileAndReturnResponse("./test.txt")
	fmt.Println(data)
}
