// Package bar is a hello world printer.
package bar

import (
	"database/sql"
	"fmt"
)

// FooService handles some stuff.
type fooService struct {
	db sql.DB
}

// Service configuration parameters.
const (
	Port    = 9090
	Timeout = 20
)

// NewService returns a configured instance of FooService.
func NewService(db sql.DB) *fooService {
	l := make([]int, 0)
	fmt.Println(l)
	return &fooService{db: db}
}
