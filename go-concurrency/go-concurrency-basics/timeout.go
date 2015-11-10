package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	BackendAPI = "http://localhost:9005?q="
	Timeout = 100
)

func main() {
	// configure http paths
	http.HandleFunc("/search", search)

	// start http server
	fmt.Println("Listening on port 9000")
	http.ListenAndServe(":9000", nil)
}

// search endpoint
func search(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	fmt.Printf("Start searching for string '%s'\n", query)
	start := time.Now()

	// infinite loop
	for {
		select {
		case res := <-searchTerm(query): // try get result
			fmt.Printf("Got response in %v.\n", time.Since(start))
			fmt.Fprint(w, string(res))
			return
		case <-time.After(time.Millisecond * Timeout): // generate message after 100ms
			fmt.Printf("Can't get response in %dms. Timeout! \n", Timeout)
			fmt.Fprint(w, "Timeout!")
			return
		}
	}
}

// searchTerm makes call to backend service for query results
func searchTerm(query string) chan []byte {
	result := make(chan []byte) // create result chan

	// run new goroutine
	go func(result chan []byte) {
		resp, err := http.Get(BackendAPI + query)
		if err != nil {
			log.Fatalf("can't make request to: %s, err: %v", BackendAPI + query, err)
		}

		if resp.Status != "200 OK" {
			log.Fatalf("wrong http status: %s", resp.Status)
		}

		bytes, _ := ioutil.ReadAll(resp.Body)

		result <- bytes
	}(result)

	// return result chan immediately (it's empty at that time)
	return result
}
