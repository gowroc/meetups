package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	as := ageService{}

	http.HandleFunc("/age", func(w http.ResponseWriter, req *http.Request) {
		request, _ := decodeAgeRequest(req)
		ageRequest := request.(calculateAgeRequest)

		age, _ := as.CalculateAge(ageRequest.YearOfBirth)

		encodeResponse(w, calculateAgeResponse{age, ""})
	})

	fmt.Printf("%v", http.ListenAndServe(":8000", nil))
}

func decodeAgeRequest(r *http.Request) (interface{}, error) {
	var request calculateAgeRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
