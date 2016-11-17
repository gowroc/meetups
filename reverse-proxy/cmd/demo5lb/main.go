package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"

	"golang.org/x/crypto/acme"
	"golang.org/x/crypto/acme/autocert"
)

var (
	prod      = flag.Bool("prod", false, "Production mode")
	httpsAddr = flag.String("https", ":443", "HTTPS address")
)

// HostSwitch maps host names to http.Handlers.
type HostSwitch map[string]http.Handler

// Implement the ServerHTTP method on our new type.
func (hs HostSwitch) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler, ok := hs[r.Host]
	if !ok {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}
	handler.ServeHTTP(w, r)
}

func newHostSwitch(reg Registry) HostSwitch {
	hs := make(HostSwitch)
	for h := range reg {
		proxy := NewMultipleHostReverseProxy(reg)
		router := http.NewServeMux()
		router.Handle("/", proxy)
		hs[h] = router
	}
	return hs
}

// Registry maps host names to slice of backends.
type Registry map[string][]string

var reg = Registry{
	"www.homedroids.io": {
		"localhost:9091",
		"localhost:9092",
	},
	"www.kowalak.cc": {
		"localhost:9093",
		"localhost:9094",
	},
}

// Hosts returns slice of host names supported by Registry.
func (r Registry) Hosts() []string {
	domains := make([]string, 0, len(reg))
	for d := range reg {
		domains = append(domains, d)
	}
	return domains
}

func NewMultipleHostReverseProxy(reg Registry) *httputil.ReverseProxy {
	return &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			req.URL.Scheme = "http"
			req.URL.Host = req.Host
		},
		Transport: &http.Transport{
			Proxy: func(req *http.Request) (*url.URL, error) {
				return http.ProxyFromEnvironment(req)
			},
			Dial: func(network, addr string) (net.Conn, error) {
				return loadBalance(network, addr, reg)
			},
			TLSHandshakeTimeout: 10 * time.Second,
		},
	}
}

func loadBalance(network, addr string, reg Registry) (net.Conn, error) {
	addr = strings.Split(addr, ":")[0]
	endpoints, ok := reg[addr]
	if !ok {
		return nil, fmt.Errorf("Not supported host: %v", addr)
	}
	for {
		// No more endpoint, stop
		if len(endpoints) == 0 {
			break
		}
		// Select a random endpoint
		i := rand.Int() % len(endpoints)
		endpoint := endpoints[i]

		// Try to connect
		conn, err := net.Dial(network, endpoint)
		if err != nil {
			// reg.Failure(serviceName, serviceVersion, endpoint, err)
			// Failure: remove the endpoint from the current list and try again.
			log.Printf("can't call endpoint %d", i)
			endpoints = append(endpoints[:i], endpoints[i+1:]...)
			continue
		}
		// Success: return the connection.
		return conn, nil
	}
	// No available endpoint.
	return nil, fmt.Errorf("No endpoint available for %s", addr)
}

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
		HostPolicy: autocert.HostWhitelist(reg.Hosts()...),
		Cache:      autocert.DirCache("certs"),
		Client:     client,
	}
	hs := newHostSwitch(reg)
	s := &http.Server{
		Addr:      *httpsAddr,
		Handler:   hs,
		TLSConfig: &tls.Config{GetCertificate: m.GetCertificate},
	}
	log.Fatal(s.ListenAndServeTLS("", ""))
}
