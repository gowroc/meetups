package handlers

import (
	"net/http"
	"bytes"
	"io"
	"fmt"
)

// START OMIT
func WithRecover(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bb := &bytes.Buffer{}
		mw := io.MultiWriter(w, bb)
		responseCopy := &responseWriter{writer: mw, responseWriter: w}
		defer func() {
			if r := recover(); r != nil{
				fmt.Printf("----\nRequest panicked.\nResponse so far\n%s\n", bb.String())
			}
		}()
		h(responseCopy, r)
	}
}
// END OMIT

