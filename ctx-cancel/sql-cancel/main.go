package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// https://github.com/go-sql-driver/mysql/issues/731
	db, err := sql.Open("mysql", "root:example@tcp(127.0.0.1:3312)/test")
	if err != nil {
		log.Fatal(err)
	}
	now := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()
	_, err = db.QueryContext(ctx, "SELECT SLEEP(10)")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Flight time %v\n", time.Since(now))
	time.Sleep(10 * time.Second)
}
