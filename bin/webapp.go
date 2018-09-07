package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	Title string
	Body  []byte
}

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "Hi, %s", request.URL.Path[1:])
}

func loadPage(title string) *Page {
	filename := title + ".html"
	body, _ := ioutil.ReadFile(filename)
	return &Page{Title: title, Body: body}
}

func editHandler(writer http.ResponseWriter, request *http.Request) {
	template, _ := template.ParseFiles("http/index.html")
	resp, _ := ioutil.ReadAll(request.Body)
	log.Printf("Request:", string(resp))
	template.Execute(writer, &Page{Title: "Opa"})
}

func main() {
	http.HandleFunc("/", editHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
