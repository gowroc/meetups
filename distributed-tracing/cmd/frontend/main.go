package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/openzipkin/zipkin-go"
	zipkinhttp "github.com/openzipkin/zipkin-go/middleware/http"

	"github.com/gowroc/meetups/distributed-tracing/lb"
	"github.com/gowroc/meetups/distributed-tracing/tracer"
)

var httpAddr = flag.String("http", ":8080", "HTTP address")

var client *zipkinhttp.Client

func backend(reg lb.Registry) http.HandlerFunc {
	type result struct {
		index int
		res   *http.Response
		err   error
	}
	return func(w http.ResponseWriter, r *http.Request) {
		resultsChan := make(chan *result, 10)

		for i := range reg {
			go func(i int) {
				log.Printf("frontend on '%s', will try backend '%s'", *httpAddr, reg[i])
				result := &result{index: i}

				span := zipkin.SpanFromContext(r.Context())
				backendReq, err := http.NewRequest("GET", fmt.Sprintf("http://%s%s", reg[i], r.URL), nil)
				if err != nil {
					result.err = err
					resultsChan <- result
					return
				}
				ctx := zipkin.NewContext(backendReq.Context(), span)
				backendReq = backendReq.WithContext(ctx)
				result.res, result.err = client.DoWithAppSpan(backendReq, "search")
				resultsChan <- result
			}(i)
		}

		var res *result

		tries := len(reg)
		for tries > 0 {
			tries--
			res = <-resultsChan
			if res.err == nil && res.res.StatusCode == http.StatusOK {
				w.WriteHeader(http.StatusOK)
				io.Copy(w, res.res.Body)
				return
			}
		}

		w.WriteHeader(http.StatusBadGateway)
	}
}

func main() {
	flag.Parse()

	reg := lb.Registry{
		"localhost:9091",
		"localhost:9092",
	}
	reg.Seed()

	trc, err := tracer.NewTracer("frontend", *httpAddr)
	if err != nil {
		log.Fatal(err)
	}

	tracingMiddleWare := zipkinhttp.NewServerMiddleware(
		trc,
		zipkinhttp.TagResponseSize(true),
	)

	client, err = zipkinhttp.NewClient(trc, zipkinhttp.ClientTrace(true))
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/", tracingMiddleWare(backend(reg)))
	log.Fatal(http.ListenAndServe(*httpAddr, nil))
}
