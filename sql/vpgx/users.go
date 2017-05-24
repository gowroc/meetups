package vpgx

import (
	"github.com/gowroc/meetups/sql/query"
	"github.com/jackc/pgx"
)

type Users struct {
	db *pgx.ConnPool
}

func NewUsers() (*Users, error) {
	pool, err := pgx.NewConnPool(pgx.ConnPoolConfig{
		ConnConfig: pgx.ConnConfig{
			Host:     "localhost",
			Port:     5433,
			User:     "postgres",
			Password: "pass",
			Database: "postgres",
		},
	})
	if err != nil {
		return nil, err
	}
	if _, err := pool.Prepare("selectAll", query.User.SelectAll); err != nil {
		return nil, err
	}
	if _, err := pool.Prepare("getByID", query.User.GetByID); err != nil {
		return nil, err
	}
	return &Users{db: pool}, err
}
