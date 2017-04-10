package handlers

import "net/http"

// START OMIT
type Decorator func(http.HandlerFunc) http.HandlerFunc

func Decorate(h http.HandlerFunc, dd ... Decorator) http.HandlerFunc {
	for _, d := range dd {
		h = d(h)
	}
	return h
}
// END OMIT