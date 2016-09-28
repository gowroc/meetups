package main

import (
	"github.com/iris-contrib/middleware/recovery"
	"github.com/kataras/iris"
)

func main() {
	iris.Use(recovery.Handler)
	iris.Get("/", func(ctx *iris.Context) {
		ctx.Write("Hi, let's panic")
		panic("Don't worry, be happy!!")
	})

	iris.Listen(":8000")
}
