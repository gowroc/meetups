package main

import (
	"flag"
	"fmt"
)

var http = flag.String("http", "", "HTTP service address (e.g., '127.0.0.1:3999')")

func main() {
	flag.Parse()
	fmt.Println(*http)
}
