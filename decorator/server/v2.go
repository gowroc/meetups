package main

import (
	"net/http"

	"github.com/gowroc/meetups/decorator/handlers"
)

// START OMIT
func addHashHandler() {
	http.HandleFunc("/hash", handlers.WithHash(handlers.Simple))
}

// END OMIT
