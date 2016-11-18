package main

import (
	"crypto/tls"
	"flag"
	"log"
	"net/http"

	"golang.org/x/crypto/acme/autocert"
)

var (
	httpsAddr = flag.String("https", ":443", "HTTPS address")
)

func main() {
	flag.Parse()
	log.Println("Starting")

	// http.ListenAndServe(":443", nil)
	m := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("www.homedroids.io"),
	}
	s := &http.Server{
		Addr:      *httpsAddr,
		TLSConfig: &tls.Config{GetCertificate: m.GetCertificate},
	}
	log.Fatal(s.ListenAndServeTLS("", ""))
}
