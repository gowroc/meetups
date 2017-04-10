package main

import (
	"net/http"

	"github.com/gowroc/meetups/decorator/handlers"
)

// START OMIT
func addGzipHashHandler() {
	http.HandleFunc("/gh", handlers.WithGzip(handlers.WithHash(handlers.Simple)))
}

// END OMIT
