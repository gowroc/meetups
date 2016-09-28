package main

import (
	"github.com/iris-contrib/middleware/basicauth"
	"github.com/kataras/iris"
)

func main() {
	authentication := basicauth.Default(map[string]string{"user": "pass"})

	iris.Get("/secret", authentication, func(ctx *iris.Context) {
		username := ctx.GetString("user")
		ctx.Write("Hello authenticated user: %s ", username)
	})

	iris.Listen(":8000")
}
