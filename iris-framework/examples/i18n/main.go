package main

import (
	"github.com/iris-contrib/middleware/i18n"
	"github.com/kataras/iris"
)

func main() {
	iris.Use(i18n.New(i18n.Config{
		Default: "en-US",
		Languages: map[string]string{
			"en-US": "./messages/en.ini",
			"pl-PL": "./messages/pl.ini",
		},
		URLParameter: "lang",
	}))

	iris.Get("/", func(ctx *iris.Context) {
		message := ctx.GetFmt("translate")("hello", "GLUG Wrocław")
		language := ctx.Get("language")

		ctx.Write("%s (in %s)", message, language)
	})

	iris.Listen(":8000")

}
