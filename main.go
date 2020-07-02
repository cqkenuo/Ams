package main

import (
	"Ams/config"
	"Ams/test"
	//"Ams/crawler"
	//"Ams/subdomain/spiders"
	"fmt"
	//"github.com/kataras/iris"
)

func init(){

}

func main() {
	test.Te()
	setting := config.LoadConfig()
	fmt.Printf("%p\n",setting)

	//f := subdomain.Factory{}
	//baiduSpider := f.CreateBaiDuSpider("oppo.com")
	//s := crawler.NewScheduler(baiduSpider,10)
	//s.Start()
	//app := iris.New()
	//app.Run(iris.Addr(":8080"))
}