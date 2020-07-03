package main

import (
	"Ams/crawler"
	"Ams/model"
	"Ams/subdomain"
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
	go subdomain.SubdomainService(subdomainChan)
}

func main() {
	Init()
	tmpCallback := make(chan []crawler.SpiderInterface)
	d := &model.Domains{Domain: "oppo.com", Fid: 1}
	subdomainChan <- subdomain.SDServiceTask{Domain: d, Callback: tmpCallback}
	for _, item := range <-tmpCallback {
		schedulerChan <- item
	}
	app := iris.New()
	app.Run(iris.Addr(":8080"))
}
