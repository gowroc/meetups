package handlers

import (
	"io"
	"net/http"
)

// START OMIT
type responseWriter struct {
	writer         io.Writer
	responseWriter http.ResponseWriter
}

func (w *responseWriter) Write(b []byte) (int, error) {
	return w.writer.Write(b)
}

func (w *responseWriter) Header() http.Header {
	return w.responseWriter.Header()
}

func (w *responseWriter) WriteHeader(i int) {
	w.responseWriter.WriteHeader(i)
}
// END OMIT
