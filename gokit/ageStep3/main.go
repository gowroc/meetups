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

	var as AgeService
	as = ageService{}
	as = loggingMiddleware{logger, as}
	ageEndpoint := makeCalculateAgeEndpoint(as)

	ageHandler := httptransport.NewServer(
		ctx,
		ageEndpoint,
		decodeAgeRequest,
		encodeResponse,
	)

	http.Handle("/age", ageHandler)
	logger.Log("msg", "HTTP", "addr", ":8003")
	logger.Log("err", http.ListenAndServe(":8003", nil))
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
