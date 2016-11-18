package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var httpAddr = flag.String("http", ":9090", "HTTP address")

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello from %s!", *httpAddr)
}

func main() {
	flag.Parse()
	log.Printf("Starting (addr=%s)", *httpAddr)
	http.HandleFunc("/", hello)
	log.Fatal(http.ListenAndServe(*httpAddr, nil))
	// server := &http.Server{Addr: *httpAddr, Handler: nil}
	// server.SetKeepAlivesEnabled(false)
	// log.Fatal(server.ListenAndServe())
}
