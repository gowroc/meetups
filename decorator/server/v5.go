package main

import (
	"net/http"

	"github.com/gowroc/meetups/decorator/handlers"
)

// START OMIT
func addGzipLoggingHashFlushHandler() {
	http.HandleFunc("/gflh",
		handlers.WithGzipFlush(
			handlers.WithLogging(
				handlers.WithHash(handlers.Simple))))
}

// END OMIT
