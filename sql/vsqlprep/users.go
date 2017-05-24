package vsqlprep

import (
	"database/sql"

	"github.com/gowroc/meetups/sql/query"
)

type Users struct {
	db        *sql.DB
	selectAll *sql.Stmt
	getByID   *sql.Stmt
}

func NewUsers() (*Users, error) {
	db, err := sql.Open("postgres",
		"user=postgres dbname=postgres password=pass port=5433 sslmode=disable")
	if err != nil {
		return nil, err
	}
	selectAll, err := db.Prepare(query.User.SelectAll)
	if err != nil {
		return nil, err
	}
	getByID, err := db.Prepare(query.User.GetByID)
	if err != nil {
		return nil, err
	}

	return &Users{
		db:        db,
		selectAll: selectAll,
		getByID:   getByID,
	}, err
}
