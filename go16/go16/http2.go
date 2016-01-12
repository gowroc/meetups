package main

import (
	"net/http"
	"log"
	"fmt"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "HTTP/2 is here!")
}

func logHttp(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		t := time.Now()
		handler.ServeHTTP(res, req)
		log.Printf("%s [%s] \"%s %s %s\" \"%s\" \"%s\" \"Took: %s\"", req.RemoteAddr,
			t.Format("02/Jan/2006:15:04:05 -0700"), req.Method, req.RequestURI, req.Proto, req.Referer(), req.UserAgent(), time.Since(t))
	})
}

func main() {
	http.HandleFunc("/", handler)
	if err := http.ListenAndServeTLS(":8090", "/home/md/cert.pem", "/home/md/key.pem", logHttp(http.DefaultServeMux)); err != nil {
		log.Fatalf("Error: %s", err)
	}
}

