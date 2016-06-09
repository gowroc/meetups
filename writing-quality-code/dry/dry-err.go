package main

import (
	"html/template"
	"net/http"
)

// Record holds record info.
type Record struct {
	ID int
}

type database struct{}

func (d *database) Get(id string, rec *Record) error {
	return nil
}

var db = new(database)

var viewTemplate = template.Must(template.New("name").Parse("html"))

func viewRecord(w http.ResponseWriter, r *http.Request) {
	rec := new(Record)
	if err := db.Get(r.FormValue("id"), rec); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	if err := viewTemplate.Execute(w, rec); err != nil {
		http.Error(w, err.Error(), 500)
	}
}

func viewRecord2(w http.ResponseWriter, r *http.Request) error {
	rec := new(Record)
	if err := db.Get(r.FormValue("id"), rec); err != nil {
		return err
	}
	return viewTemplate.Execute(w, rec)
}

type appHandler func(http.ResponseWriter, *http.Request) error

func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := fn(w, r); err != nil {
		http.Error(w, err.Error(), 500)
	}
}

func main() {
	http.HandleFunc("/view", viewRecord)
	http.Handle("/view2", appHandler(viewRecord2))
	http.ListenAndServe(":8080", nil)
}
