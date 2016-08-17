package main

import (
	"errors"
	"time"
)

type ageService struct{}

func (ageService) CalculateAge(yearOfBirth int) (int, error) {
	year := time.Now().Year()

	if yearOfBirth > year {
		return 0, errNotBornYet
	}

	return year - yearOfBirth, nil
}

var errNotBornYet = errors.New("Not born yet")

type calculateAgeRequest struct {
	YearOfBirth int `json:"yearOfBirth"`
}

type calculateAgeResponse struct {
	Age int    `json:"age"`
	Err string `json:"err,omitempty"`
}
