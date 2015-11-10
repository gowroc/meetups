package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var backendServers = map[int]string{
	0: "http://localhost:9001?q=",
	1: "http://localhost:9002?q=",
	2: "http://localhost:9003?q=",
}

func main() {
	http.HandleFunc("/search", searchFor)
	fmt.Println("Listening on port 9000")
	http.ListenAndServe(":9000", nil)
}

func searchFor(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	fmt.Printf("Start searching for string=%s\n", query)

	resultChan := make(chan response)
	for i := 0; i < 3; i++ {
		go searchInBackend(i, query, resultChan)
	}

	res := <-resultChan
	fmt.Printf("Got response from server with id: %d\n", res.serverId)
	fmt.Fprint(w, string(res.payload))
}

func searchInBackend(id int, query string, result chan response) {
	resp, err := http.Get(backendServers[id] + query)
	if err != nil {
		log.Printf("can't make request to: %s, err: %v", backendServers[id]+query, err)
		return
	}

	if resp.Status != "200 OK" {
		log.Printf("wrong http status: %s", resp.Status)
		return
	}

	bytes, _ := ioutil.ReadAll(resp.Body)

	result <- response{id, bytes}
}

type response struct {
	serverId int
	payload  []byte
}
