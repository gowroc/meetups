package main

import (
	"net/http"

	"github.com/gowroc/meetups/decorator/handlers"
)

// START OMIT
func addSimpleHandler() {
	http.HandleFunc("/simple", handlers.Simple)
}

// END OMIT
