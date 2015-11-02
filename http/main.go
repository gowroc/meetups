package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// Weather information for location.
type Weather struct {
	Temp  int  `json:"temp"`
	Windy bool `json:"windy"`
}

var cityWeather = map[string]Weather{}

func weather(rw http.ResponseWriter, req *http.Request) {
	city := req.URL.Path[len("/weather/"):]
	switch req.Method {
	case "GET":
		resp, err := json.Marshal(cityWeather[city])
		if err != nil {
			http.Error(rw, "Internal server error", http.StatusInternalServerError)
			return
		}
		rw.Write(resp)
	case "POST":
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Printf("can't read request body: %v", err)
			http.Error(rw, "Internal server error", http.StatusInternalServerError)
			return
		}
		var w Weather
		if err := json.Unmarshal(body, &w); err != nil {
			log.Printf("can't decode request body: %v", err)
			http.Error(rw, "Bad request", http.StatusBadRequest)
			return
		}
		cityWeather[city] = w
	default:
		http.Error(rw, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	log.Println("Starting service")
	http.HandleFunc("/weather/", weather)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
