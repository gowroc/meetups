package main

import (
	"github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"
)

func makeCalculateAgeEndpoint(as AgeService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(calculateAgeRequest)
		age, err := as.CalculateAge(req.YearOfBirth)
		if err != nil {
			return calculateAgeResponse{age, err.Error()}, nil
		}
		return calculateAgeResponse{age, ""}, nil
	}
}
