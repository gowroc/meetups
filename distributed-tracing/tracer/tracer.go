package tracer

import (
	"fmt"

	"github.com/openzipkin/zipkin-go"
	reporterhttp "github.com/openzipkin/zipkin-go/reporter/http"
)

const zipkinURL = "http://localhost:9411/api/v2/spans"

func NewTracer(serviceName string, hostPort string) (*zipkin.Tracer, error) {
	// The reporter sends traces to Zipkin server.
	reporter := reporterhttp.NewReporter(zipkinURL)

	// Local endpoint represent the local service information.
	localEndpoint, err := zipkin.NewEndpoint(serviceName, hostPort)
	if err != nil {
		return nil, fmt.Errorf("cannot create zipkin local endpoint: %v", err)
	}

	// Sampler tells you which traces are going to be sampled or not. In this case we will record 100% (1.00) of traces.
	sampler, err := zipkin.NewCountingSampler(1)
	if err != nil {
		return nil, fmt.Errorf("cannot create zipkin sampler: %v", err)
	}

	t, err := zipkin.NewTracer(
		reporter,
		zipkin.WithSampler(sampler),
		zipkin.WithLocalEndpoint(localEndpoint),
	)
	if err != nil {
		return nil, fmt.Errorf("cannot create zipkin tracer: %v", err)
	}

	return t, err
}
