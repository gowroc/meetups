package main

import (
	"encoding/json"
	"log"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"golang.org/x/net/context"
)

func main() {
	ctx := context.Background()
	as := ageService{}

	ageHandler := httptransport.NewServer(
		ctx,
		makeCalculateAgeEndpoint(as),
		decodeAgeRequest,
		encodeResponse,
	)

	http.Handle("/age", ageHandler)
	log.Fatal(http.ListenAndServe(":8001", nil))
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
