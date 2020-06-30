package Subdomain

import (
	"Ams/Crawler"
	"net/http"
)

type BaiDuSpider struct {
	Crawler.Spider
}

func (s *BaiDuSpider)Parse(request *http.Request, response *Crawler.CResponse) Crawler.SpiderResult{
	return Crawler.SpiderResult{}
}

func (s *BaiDuSpider)Seeds() []*Crawler.Task{
	return nil
}


func (s *BaiDuSpider)ResultProcess(result []map[string]interface{}){

}