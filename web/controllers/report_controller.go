package controllers

import (
	"Ams/crawler"
	"Ams/subdomain"
	"Ams/web/services"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type ReportController struct {
	Ctx           *iris.Context
	Service       services.AmsServicesInterFace
	SubdomainChan chan subdomain.SDServiceTask
	SchedulerChan chan crawler.SpiderInterface
}

func (c *ReportController) Get() mvc.Result {
	return mvc.Response{}
}

// 添加根域名
func (c *ReportController) PostAddRootDomain() mvc.Result {
	return mvc.Response{}
}
// 导入域名
func (c *ReportController) PostImportDomains() mvc.Result {
	return mvc.Response{}
}
// 单独采集信息
func (c *ReportController) PostCollectInfo() mvc.Result {
	return mvc.Response{}
}
// 下一页
func (c *ReportController) PostNextPage() mvc.Result {
	return mvc.Response{}
}
