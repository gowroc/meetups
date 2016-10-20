package main

import "github.com/kataras/iris"

func main() {
	iris.StaticWeb("/", "./public", 0)
	iris.Listen(":8000")
}
