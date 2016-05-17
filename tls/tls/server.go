package main

import (
	"net/http"
	"log"
	"os"
)

func handler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Hello HTTP/2!\n"))
}

func main() {
	if len(os.Args) != 4 {
		log.Fatal("program arguments: port cert key")
	}
	port, cert, key := os.Args[1], os.Args[2], os.Args[3]
	http.HandleFunc("/", handler)
	err := http.ListenAndServeTLS(":" + port, cert, key, nil)
	log.Fatal(err)
}
