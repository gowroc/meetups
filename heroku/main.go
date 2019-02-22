package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(rw, "Hello GoWroc!")
	})
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
