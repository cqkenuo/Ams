package main

import (
	"Ams/crawler"
	"Ams/subdomain"
	"Ams/web"
	"github.com/kataras/iris"
)

var (
	subdomainChan chan subdomain.SDServiceTask
	schedulerChan chan crawler.SpiderInterface
)

func Init() {
	subdomainChan = make(chan subdomain.SDServiceTask)
	schedulerChan = make(chan crawler.SpiderInterface)
	go crawler.SchedulerService(schedulerChan)
	go subdomain.Service(subdomainChan)
}

func main() {
	Init()
	//tmpCallback := make(chan []crawler.SpiderInterface)
	//d := &model.Domains{Domain: "oppo.com", Fid: 0}
	//subdomainChan <- subdomain.SDServiceTask{Domain: d, Callback: tmpCallback}
	//for _, item := range <-tmpCallback {
	//	schedulerChan <- item
	//}
	app := web.NewWebApp(subdomainChan, schedulerChan)
	app.Run(iris.Addr(":8080"))
}
