package main

import (
	"fmt"
	"net/http"
)

func main() {
	addSimpleHandler()
	addHashHandler()
	addGzipHashHandler()
	addGzipLoggingHashHandler()
	addGzipLoggingHashFlushHandler()
	addDecoratorsHandler()
	fmt.Println(http.ListenAndServe(":8888", nil))
}
