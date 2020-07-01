package subdomain

import (
	"Ams/crawler"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type BaiDuSpider struct {
	baseSpider
}

func(s *BaiDuSpider)getUrl(query string,pageNumber int) string {
	return fmt.Sprintf("%s?pn=%d&wd=%s&oq=%s",s.baseUrl,pageNumber,query,query)
}

func(s *BaiDuSpider)getQuery(domain string,exceptDomain []string)string{
	if exceptDomain != nil && len(exceptDomain) >= 2{
		exceptDomain = exceptDomain[:2]
		return url.QueryEscape(fmt.Sprintf("site:%s -site:www.%s -site:%s",domain,domain,strings.Join(exceptDomain," -site:")))
	}
	return url.QueryEscape(fmt.Sprintf("site:%s -site:www.%s",domain,domain))
}

func (s *BaiDuSpider)Parse(task *crawler.Task, response *crawler.CResponse) crawler.SpiderResult{
	fmt.Println(task.GetAttach())
	return crawler.SpiderResult{}
}

func (s *BaiDuSpider)Seeds() []*crawler.Task{
	var seeds []*crawler.Task
	for i:=0;i<=170;i+=10{
		u := s.getUrl(s.getQuery(s.domain,[]string{}),i)
		r,_ := http.NewRequest("GET",u,nil)
		t := crawler.NewTask(r,nil)
		fmt.Println(u)
		t.SetAttach(map[string]interface{}{
			"pageNumber":i,
			"domain":s.domain,
		})
		seeds = append(seeds, t)
	}
	return seeds
}


func (s *BaiDuSpider)ResultProcess(result []map[string]interface{}){

}