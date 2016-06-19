package foo

import (
	"database/sql"
)

// START OMIT
type FooService struct {
	db sql.DB
}

func NewService(db sql.DB) *FooService {
	return &FooService{db: db}
}

// END OMIT
