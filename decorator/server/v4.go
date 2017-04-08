package main

import (
	"net/http"

	"github.com/gowroc/meetups/decorator/handlers"
)

// START OMIT
func addGzipLoggingHashHandler() {
	http.HandleFunc("/glh",
		handlers.WithGzip(
			handlers.WithLogging(
				handlers.WithHash(
					handlers.Simple))))
}

// END OMIT
