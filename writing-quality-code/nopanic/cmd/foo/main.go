package main

import (
	"database/sql"
	"log"

	"github.com/gowroc/meetups/writing-quality-code/nopanic"
)

func main() {
	db, err := sql.Open("postgres", config.PSQLConn)
	if err != nil {
		log.Fatal("can't open database conn:", err)
	}
	fs := foo.NewService(db)
}
