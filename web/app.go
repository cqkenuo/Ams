package web

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

func NewWebApp () *iris.Application{
	app := iris.New()

	mvc.New(app.Party("/"))

	return app
}