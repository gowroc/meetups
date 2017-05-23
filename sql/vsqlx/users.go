package vsqlx

import "github.com/jmoiron/sqlx"

type Users struct {
	db *sqlx.DB
}

func NewUsers() (*Users, error) {
	db, err := sqlx.Open("postgres",
		"user=postgres dbname=postgres password=pass port=5433 sslmode=disable")
	return &Users{db: db}, err
}
