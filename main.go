package main

import (
	"Ams/crawler"
	subdomain "Ams/subdomain/spiders"
	"github.com/kataras/iris"
)

func init(){

}

func main() {
	//test.Te(map[string]int{"v":10})
	f := subdomain.Factory{}
	baiduSpider := f.CreateBaiDuSpider("oppo.com")
	s := crawler.NewScheduler(baiduSpider,10)
	s.Start()
	app := iris.New()
	app.Run(iris.Addr(":8080"))
}