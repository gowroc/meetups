package handlers

import (
	"compress/gzip"
	"net/http"
)

// START OMIT
func WithGzip(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Encoding", "gzip")
		gz := gzip.NewWriter(w)
		defer gz.Close()
		gzr := responseWriter{writer: gz, responseWriter: w}
		fn(&gzr, r)
	}
}

// END OMIT
