package main

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/kataras/iris"
	_ "github.com/lib/pq"

	"github.com/gowroc/meetups/iris-framework/examples/crud/posts"
)

func main() {
	dbinfo := fmt.Sprintf("port=32768 user=%s password=%s dbname=%s sslmode=disable", "postgres", "pass123", "postgres")
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		fmt.Printf("err: %+v ", err)
	}
	defer db.Close()

	iris.Get("/people", GetAll(db))
	iris.Get("/people/:id", Get(db))
	iris.Post("/people/:id", Post(db))
	iris.Delete("/people/:id", Delete(db))

	iris.Listen(":8000")
}

func GetAll(db *sql.DB) iris.HandlerFunc {
	return func(ctx *iris.Context) {
		p, err := posts.GetAll(db)
		if err != nil {
			fmt.Printf("%v", err)
			ctx.Error("Failed to load people list", 503)
		}

		ctx.JSON(200, p)
	}
}

func Get(db *sql.DB) iris.HandlerFunc {
	return func(ctx *iris.Context) {
		ID, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			handleError(ctx, "Failed to load people list", err)
		}

		p, err := posts.Get(db, ID)
		if err != nil {
			handleError(ctx, "Failed to load people list", err)
		}

		ctx.JSON(200, p)
	}
}

func Post(db *sql.DB) iris.HandlerFunc {
	return func(ctx *iris.Context) {
		name := ctx.Param("name")
		hobby := ctx.Param("hobby")

		err := posts.Post(db, name, hobby)
		if err != nil {
			handleError(ctx, "Failed to save person", err)
			return
		}

		ctx.Write("ok")
	}
}

func Delete(db *sql.DB) iris.HandlerFunc {
	return func(ctx *iris.Context) {
		ID, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			handleError(ctx, "Failed to load people list", err)
		}

		err = posts.Delete(db, ID)
		if err != nil {
			handleError(ctx, "Failed to load people list", err)
		}

		ctx.Write("ok")
	}
}

func handleError(ctx *iris.Context, message string, err error) {
	fmt.Printf("%v", err)
	ctx.Error(message, 503)
}
