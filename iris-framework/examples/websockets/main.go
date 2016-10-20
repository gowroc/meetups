package main

import (
	"fmt"

	"github.com/kataras/iris"
)

type clientPage struct {
	Title string
	Host  string
}

func main() {
	iris.Static("/js", "./static/js", 1)

	iris.Get("/", func(ctx *iris.Context) {
		ctx.Render("client.html", clientPage{"Client Page", ctx.HostString()})
	})

	iris.Config.Websocket.Endpoint = "/gowroc"

	iris.Websocket.OnConnection(func(c iris.WebsocketConnection) {
		c.Join("room")
		c.On("message", func(message string) {
			c.To("room").Emit("message", "From: "+c.ID()+": "+message)
		})

		c.OnDisconnect(func() {
			fmt.Printf("\nConnection with ID: %s has been disconnected!", c.ID())
		})
	})

	iris.Listen(":8000")
}
