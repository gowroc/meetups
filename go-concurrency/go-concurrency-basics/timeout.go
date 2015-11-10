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
	http.HandleFunc("/search", search)
	fmt.Println("Listening on port 9000")
	http.ListenAndServe(":9000", nil)
}

func search(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	fmt.Printf("Start searching for string '%s'\n", query)

	result := searchTerm(query)

	for {
		select {
		case res := <-result:
			fmt.Fprint(w, string(res))
			return
		case <-time.After(time.Millisecond * Timeout):
			fmt.Printf("Can't get response in %dms. Timeout! \n", Timeout)
			fmt.Fprint(w, "Timeout!")
			return
		}
	}
}

func searchTerm(query string) chan []byte {
	result := make(chan []byte)
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
	return result
}
