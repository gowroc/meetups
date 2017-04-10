package handlers

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

// START OMIT
func WithLogging(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bb := &bytes.Buffer{}
		mw := io.MultiWriter(w, bb)
		responseCopy := &responseWriter{writer: mw, responseWriter: w}
		h(responseCopy, r)
		fmt.Printf("\n----\nRequest finished.\nResponse\n%s\n", bb.String())
	}
}

// END OMIT
