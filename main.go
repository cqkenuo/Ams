package main

import (
	"Ams/config"
	//"Ams/crawler"
	//"Ams/subdomain/spiders"
	"fmt"
	//"github.com/kataras/iris"
)
func main() {
	setting := config.LoadConfig()
	fmt.Println(setting.DbConf.DbCharSet)
	//f := subdomain.Factory{}
	//baiduSpider := f.CreateBaiDuSpider("oppo.com")
	//s := crawler.NewScheduler(baiduSpider,10)
	//s.Start()
	//app := iris.New()
	//app.Run(iris.Addr(":8080"))
}