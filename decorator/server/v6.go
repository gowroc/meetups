package main

import (
	"net/http"

	"github.com/gowroc/meetups/decorator/handlers"
)

// START OMIT
func addDecoratorsHandler() {
	http.HandleFunc("/decorated",
		handlers.Decorate(
			handlers.Simple,
			handlers.WithHash,
			handlers.WithLogging,
			handlers.WithGzip,
		))

}

// END OMIT
