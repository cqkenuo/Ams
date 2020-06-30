package main

import (
	"Ams/Crawler"
	"fmt"
	"github.com/kataras/iris"
	"net/http"
)

type BaiDuSpider struct {
	Crawler.Spider
}

func (s *BaiDuSpider)Seeds() []*Crawler.Task {
	req,_ := http.NewRequest("POST","http://www.baidu.com",nil)
	var ts []*Crawler.Task
	ts = append(ts,Crawler.NewTask(req,s.v))
	fmt.Println("到这2")
	return ts
}

func (s *BaiDuSpider)v(r *http.Request,resp *Crawler.CResponse)  Crawler.SpiderResult {
	var r1 []map[string]interface{}
	r1 = append(r1,map[string]interface{}{"zi":1})
	return Crawler.SpiderResult{ResultType: Crawler.ResultType, SetData: r1}
}

func (s *BaiDuSpider)Parse(request *http.Request, response *Crawler.CResponse) Crawler.SpiderResult{
	return Crawler.SpiderResult{}
}

func (s *BaiDuSpider)ResultProcess(result []map[string]interface{}){
	fmt.Println(result)
}

func main() {
	b := &BaiDuSpider{}
	s := Crawler.NewScheduler(b,1)
	s.Start()
	app := iris.New()
	app.Run(iris.Addr(":8080"))
}