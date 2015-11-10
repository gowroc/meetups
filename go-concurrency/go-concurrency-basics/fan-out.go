package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// list of backend workers
var backendServers = map[int]string{
	0: "http://localhost:9001?q=",
	1: "http://localhost:9002?q=",
	2: "http://localhost:9003?q=",
}

func main() {
	// configure http paths
	http.HandleFunc("/search", searchFor)

	// start http server
	fmt.Println("Listening on port 9000")
	http.ListenAndServe(":9000", nil)
}

// endpoint which will be triggered when user create request for /search path
func searchFor(w http.ResponseWriter, r *http.Request) {
	// get search term from url query string
	query := r.URL.Query().Get("q")
	fmt.Printf("Start searching for string=%s\n", query)

	// create a channel for responses
	resultChan := make(chan response)

	// run queries to all backend services
	for k := range backendServers {
		go func() {
			resultChan <- searchInBackend(k, query)
		}()
	}

	// get only first result and write answer to response
	res := <-resultChan
	fmt.Printf("Got response from server with id: %d\n", res.serverId)
	fmt.Fprint(w, string(res.payload))
}

// make a query to a backend service
func searchInBackend(id int, query string) response {
	// make GET request
	resp, err := http.Get(backendServers[id] + query)
	if err != nil {
		log.Fatalf("can't make request to: %s, err: %v", backendServers[id]+query, err)
	}

	if resp.Status != "200 OK" {
		log.Fatalf("wrong http status: %s", resp.Status)
	}

	// read response and return bytes
	bytes, _ := ioutil.ReadAll(resp.Body)
	return response{id, bytes}
}

type response struct {
	serverId int
	payload  []byte
}
