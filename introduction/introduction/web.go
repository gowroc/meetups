package main

import (
	"html/template"
	"net/http"
)

var tmpl = template.Must(template.ParseFiles("web.html"))

func hello(writer http.ResponseWriter, r *http.Request) {
	tmpl.Execute(writer, "Hello World!")
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8080", nil)
}
