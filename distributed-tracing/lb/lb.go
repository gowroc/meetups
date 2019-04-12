package lb

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"

	"github.com/openzipkin/zipkin-go"
	zipkinhttp "github.com/openzipkin/zipkin-go/middleware/http"
)

// Registry holds a list of backends.
type Registry []string

func (r Registry) Seed() {
	rand.Seed(time.Now().Unix())
}

func isExcluded(endpoint string, excluded []string) bool {
	for _, e := range excluded {
		if e == endpoint {
			return true
		}
	}
	return false
}

// Roll a dice to select a healthy backend.
func (r Registry) Roll(excluded ...string) (string, error) {
	for {
		if len(excluded) >= len(r) {
			// No more endpoints to try.
			return "", errors.New("no more endpoints to try")
		}

		// Roll a dice as long as we find healthy endpoint.
		i := rand.Int() % len(r)
		if isExcluded(r[i], excluded) {
			continue
		}
		log.Printf("got endpoint %s\n", r[i])
		return r[i], nil
	}
}

// NewLoadBalancingReverseProxy returns a new ReverseProxy with load balancing.
func NewLoadBalancingReverseProxy(reg Registry, trc *zipkin.Tracer) *httputil.ReverseProxy {
	zipkinTransport, err := zipkinhttp.NewTransport(
		trc,
		zipkinhttp.TransportTrace(true),
		zipkinhttp.RoundTripper(&http.Transport{
			Proxy: func(req *http.Request) (*url.URL, error) {
				return http.ProxyFromEnvironment(req)
			},
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				return loadBalance(network, addr, reg)
			},
			DisableKeepAlives: true,
		}),
	)
	if err != nil {
		log.Fatal(err)
	}
	return &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			req.URL.Scheme = "http"
			req.URL.Host = req.Host
		},
		Transport: zipkinTransport,
	}
}

func loadBalance(network, addr string, reg Registry) (net.Conn, error) {
	addr = strings.Split(addr, ":")[0]
	var excluded []string
	for {
		endpoint, err := reg.Roll(excluded...)
		if err != nil {
			// No available endpoint.
			return nil, fmt.Errorf("no endpoint available for %s", addr)
		}

		conn, err := net.Dial(network, endpoint)
		if err != nil {
			fmt.Printf("can't connect to endpoint: %s\n", endpoint)
			excluded = append(excluded, endpoint)
			continue
		}
		// Success: return the connection.
		return conn, nil
	}
}
