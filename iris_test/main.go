// iris_test project main.go
package main

import (
	"github.com/kataras/iris"
)

func main() {

	// app := iris.Default()
	app := iris.New()
	app.Get("/", func(ctx iris.Context) {
		ctx.HTML("<b>Hello!</b>")
	})

	app.Get("/ping", func(ctx iris.Context) {
		ctx.JSON(iris.Map{
			"message": "Pong",
		})
	})

	app.Get("/test", func(ctx iris.Context) {
		ctx.XML(iris.Map{
			"message": "Pong",
		})
	})

	app.Get("/test1", func(ctx iris.Context) {
		ctx.WriteString("ok")
	})

	app.Run(iris.Addr(":8080"))
}
