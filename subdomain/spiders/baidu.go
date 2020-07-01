package subdomain

import (
	"Ams/crawler"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/jpillora/go-tld"
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
	var result []map[string]interface{}
	response.Extract.Find("a[class*='c-showurl']").Each(func(i int, selection *goquery.Selection) {
		v := selection.Text()
		if !strings.HasPrefix(v,"http") {
			v = fmt.Sprintf("http://%s",v)
		}
		u,err := tld.Parse(v)
		if err == nil{
			result = append(result,map[string]interface{}{
				"domain":u.Host,
			})
		}
	})
	return crawler.SpiderResult{SetData: result}
}

func (s *BaiDuSpider)Seeds() []*crawler.Task{
	var seeds []*crawler.Task
	for i:=0;i<=170;i+=10{
		u := s.getUrl(s.getQuery(s.domain,[]string{}),i)
		r,_ := http.NewRequest("GET",u,nil)
		t := s.NewTask(r,nil)
		t.SetAttach(map[string]interface{}{
			"pageNumber":i,
			"domain":s.domain,
		})
		seeds = append(seeds, t)
	}
	return seeds
}

func (s *BaiDuSpider)ResultProcess(result []map[string]interface{}){
	fmt.Println(result)
}