package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Weather information for location.
type Weather struct {
	Temp  int  `json:"temp"`
	Windy bool `json:"windy"`
}

var cityWeather = map[string]Weather{}

func weather(rw http.ResponseWriter, req *http.Request) {
	city := mux.Vars(req)["city"]
	resp, err := json.Marshal(cityWeather[city])
	if err != nil {
		http.Error(rw, "Internal server error", http.StatusInternalServerError)
		return
	}
	rw.Write(resp)

}

func setWeather(rw http.ResponseWriter, req *http.Request) {
	city := mux.Vars(req)["city"]
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(rw, "Internal server error", http.StatusInternalServerError)
		return
	}
	var w Weather
	if err := json.Unmarshal(body, &w); err != nil {
		http.Error(rw, "Bad request", http.StatusBadRequest)
		return
	}
	cityWeather[city] = w
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/weather/{city}/", weather).Methods("GET")
	r.HandleFunc("/weather/{city}/", setWeather).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", r))
}
