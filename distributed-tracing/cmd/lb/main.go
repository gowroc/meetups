package main

import (
	"flag"
	"log"
	"net/http"

	zipkinhttp "github.com/openzipkin/zipkin-go/middleware/http"

	"github.com/gowroc/meetups/distributed-tracing/lb"
	"github.com/gowroc/meetups/distributed-tracing/tracer"
)

var (
	httpAddr = flag.String("http", "localhost:8000", "HTTP address")
)

func main() {
	flag.Parse()

	trc, err := tracer.NewTracer("lb", *httpAddr)
	if err != nil {
		log.Fatal(err)
	}
	zipkinMiddleWare := zipkinhttp.NewServerMiddleware(
		trc,
		zipkinhttp.TagResponseSize(true),
	)

	reg := lb.Registry{
		"localhost:8081",
		"localhost:8082",
	}
	reg.Seed()
	balancer := lb.NewLoadBalancingReverseProxy(reg, trc)
	log.Fatal(http.ListenAndServe(*httpAddr, zipkinMiddleWare(balancer)))
}
