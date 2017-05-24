package vsql

import "database/sql"

type Users struct {
	db *sql.DB
}

func NewUsers() (*Users, error) {
	db, err := sql.Open("postgres",
		"user=postgres dbname=postgres password=pass port=5433 sslmode=disable")
	return &Users{db: db}, err
}
