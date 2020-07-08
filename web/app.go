package web

import (
	"Ams/crawler"
	"Ams/subdomain"
	"Ams/web/controllers"
	"Ams/web/services"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

func NewWebApp(subdomainChan chan subdomain.SDServiceTask, schedulerChan chan crawler.SpiderInterface) *iris.Application {
	app := iris.New()
	report := mvc.New(app.Party("/report"))
	report.Register(services.NewWebService(), subdomainChan, schedulerChan)
	report.Handle(new(controllers.ReportController))
	return app
}
