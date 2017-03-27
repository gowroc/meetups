package db

import (
	"database/sql"
	"fmt"

	// necessary for postgres driver to work within sql package
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

// User represents user details in SQL database.
type User struct {
	Login string `json:"login"`
	Admin string `json:"admin"`
}

// Postgres represents an implementation of SQL database.
type Postgres struct {
	conn *sql.DB
}

// NewPostgres returns an instance of SQL database.
func NewPostgres(host string, port int, user, dbName string) (*Postgres, error) {
	dataSourceName := fmt.Sprintf("user=%s dbname=%s sslmode=disable port=%d", user, dbName, port)
	conn, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create db connection")
	}
	if err := conn.Ping(); err != nil {
		return nil, errors.Wrap(err, "failed to connect to database")
	}

	return &Postgres{
		conn: conn,
	}, nil
}

// Close closes SQL connection.
func (p Postgres) Close() error {
	return p.conn.Close()
}

// GetOne returns single (random) user from the database.
func (p Postgres) GetOne() (*User, error) {
	rows, err := p.conn.Query("SELECT * from users LIMIT 1")
	if err != nil {
		return nil, errors.Wrap(err, "failed to read from database")
	}
	var u User
	exist := rows.Next()
	if exist {
		if err = rows.Scan(&u.Login, &u.Admin); err != nil {
			return nil, errors.Wrap(err, "failed to load user data")
		}
	}

	return &u, nil
}

// GetUser returns an user with given username.
func (p Postgres) GetUser(username string) (*User, error) {
	rows, err := p.conn.Query("SELECT * from users WHERE username = $1", username)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read from database")
	}
	var u User
	exist := rows.Next()
	if exist {
		if err = rows.Scan(&u.Login, &u.Admin); err != nil {
			return nil, errors.Wrap(err, "failed to load user data")
		}
	}

	return &u, nil
}

/*
psql, err := db.NewPostgres("localhost", 15432, "postgres", "graphql")
if err != nil {
	log.Fatalf("failed to connect to Postgres: %v", err)
}
defer psql.Close()
*/
