package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func weather(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("weather is good"))
}

func setWeather(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("weather is set"))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/weather/{city}/", weather).Methods("GET")
	r.HandleFunc("/weather/{city}/", setWeather).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", r))
}
