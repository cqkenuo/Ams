package main

import (
	"Ams/crawler"
	"Ams/subdomain/spiders"
	"github.com/kataras/iris"
)
func main() {
	f := subdomain.Factory{}
	baiduSpider := f.CreateBaiDuSpider("oppo.com")
	s := crawler.NewScheduler(baiduSpider,10)
	s.Start()
	app := iris.New()
	app.Run(iris.Addr(":8080"))
}