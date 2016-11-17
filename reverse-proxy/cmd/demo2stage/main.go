package main

import (
	"crypto/tls"
	"flag"
	"log"
	"net/http"

	"golang.org/x/crypto/acme"
	"golang.org/x/crypto/acme/autocert"
)

var (
	prod      = flag.Bool("prod", false, "Production mode")
	httpsAddr = flag.String("https", ":443", "HTTPS address")
)

func main() {
	flag.Parse()
	log.Printf("Starting (prod=%t)", *prod)

	client := &acme.Client{}
	if !*prod {
		client = &acme.Client{
			DirectoryURL: "https://acme-staging.api.letsencrypt.org/directory",
		}
	}
	m := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("www.homedroids.io"),
		Client:     client,
	}
	s := &http.Server{
		Addr:      *httpsAddr,
		TLSConfig: &tls.Config{GetCertificate: m.GetCertificate},
	}
	log.Fatal(s.ListenAndServeTLS("", ""))
}
