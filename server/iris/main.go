package main

import (
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()
	app.Use(iris.Compression)

	app.Get("/", func(ctx iris.Context) {
		_ = ctx.JSON(iris.Map{
			"message": "pong",
		})
	})

	_ = app.Listen("0.0.0.0:8080")
}
