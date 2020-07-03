package main

import (
	"Ams/crawler"
	"Ams/model"
	subdomain "Ams/subdomain/spiders"
	"github.com/kataras/iris"
)

func init() {

}

func main() {
	//test.Te(map[string]int{"v":10})
	f := subdomain.Factory{}
	s := f.CreateSpider(&model.Domains{Domain: "oppo.com", Fid: 1})
	for _, spider := range s {
		c := crawler.NewScheduler(spider, 10)
		c.Start()
	}
	//baiduSpider := f.CreateBaiDuSpider("oppo.com")
	//s := crawler.NewScheduler(baiduSpider,10)
	//s.Start()
	app := iris.New()
	app.Run(iris.Addr(":8080"))
}
