package subdomain

import (
	"Ams/crawler"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/jpillora/go-tld"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"sync"
)

var lock = sync.Mutex{}

type BaiDuSpider struct {
	baseSpider
	domains map[string]int
}

type DomainCount struct {
	Domain string
	Count int
}
type AllDomain []DomainCount
func (p AllDomain) Swap(i, j int) {p[i], p[j] = p[j], p[i]}
func (p AllDomain) Len() int{return len(p)}
func (p AllDomain) Less(i, j int) bool { return p[i].Count > p[j].Count }

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
	var tasks []*crawler.Task
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

	if len(s.domains) >= 2{
		p := make(AllDomain, len(s.domains))
		i := 0
		for k,v := range s.domains{
			p[i] = DomainCount{k,v}
			i ++
		}
		sort.Sort(p)
		except := make([]string,2)
		for _,item := range p[:2]{
			except = append(except,item.Domain)
		}
		page,o := task.GetAttach()["pageNumber"].(int)
		if o{
			r,_ := http.NewRequest("GET",s.getUrl(s.getQuery(s.domain,except),page),nil)
			t := s.NewTask(r,nil)
			t.SetAttach(map[string]interface{}{
				"pageNumber":page,
				"domain":s.domain,
			})
			tasks = append(tasks,t)
		}
	}
	return crawler.SpiderResult{SetData: result,TaskData: tasks}
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
	lock.Lock()
	for _,item := range result{
		domain,ok1 := item["domain"].(string)
		if ok1 {
			v,ok := s.domains[domain]
			if ok {
				s.domains[domain] = v + 1
			}else {
				s.domains[domain] = 1
			}
		}
	}
	lock.Unlock()
	// 保存逻辑
}

func (s *BaiDuSpider)Close(){

}