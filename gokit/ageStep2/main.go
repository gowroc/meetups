package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/go-kit/kit/log"

	httptransport "github.com/go-kit/kit/transport/http"
	"golang.org/x/net/context"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)

	ctx := context.Background()

	as := ageService{}
	// var ageEndpoint endpoint.Endpoint
	ageEndpoint := makeCalculateAgeEndpoint(as)
	ageEndpoint = loggingMiddleware(log.NewContext(logger).With("method", "CalculateAge"))(ageEndpoint)

	ageHandler := httptransport.NewServer(
		ctx,
		ageEndpoint,
		decodeAgeRequest,
		encodeResponse,
	)

	http.Handle("/age", ageHandler)
	logger.Log("msg", "HTTP", "addr", ":8002")
	logger.Log("err", http.ListenAndServe(":8002", nil))
}

func decodeAgeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request calculateAgeRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
