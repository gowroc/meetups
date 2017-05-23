package vmodl

import (
	"database/sql"

	"github.com/gowroc/meetups/sql/user"
	"github.com/jmoiron/modl"
)

type Users struct {
	db *modl.DbMap
}

func NewUsers() (*Users, error) {
	db, err := sql.Open("postgres",
		"user=postgres dbname=postgres password=pass port=5433 sslmode=disable")
	if err != nil {
		return nil, err
	}
	m := modl.NewDbMap(db, modl.PostgresDialect{})
	m.AddTable(user.User{}, "users").SetKeys(false, "user_id")
	m.AddTable(user.Post{}, "user_posts").SetKeys(false, "user_post_id")
	return &Users{db: m}, nil
}
