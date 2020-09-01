package main

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

func main() {
	var b bytes.Buffer
	writer := multipart.NewWriter(&b)
	if err := writer.WriteField("title", "golang"); err != nil {
		log.Panicln(err)
	}
	fileWriter, err := writer.CreateFormFile("thumbnail", "golang.jpg")
	if err != nil {
		log.Panicln(err)
	}
	readFile, err := os.Open("golang.jpg")
	if err != nil {
		log.Panicln(err)
	}
	defer readFile.Close()

	_, err = io.Copy(fileWriter, readFile)
	if err != nil {
		log.Panicln(err)
	}
	writer.Close()

	res, err := http.Post("https://yourdestination", writer.FormDataContentType(), &b)
	if err != nil {
		log.Panicln(err)
	}
	log.Println("response status: ", res.Status)
}
