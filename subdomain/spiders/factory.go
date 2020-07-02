package subdomain

import (
	"Ams/crawler"
	"fmt"
	"github.com/jpillora/go-tld"
	"strings"
)

type Factory struct {

}

func (f *Factory)validDomain(domain string)(*tld.URL,error){
	if !strings.HasPrefix(domain,"http://") && !strings.HasPrefix(domain,"https://"){
		domain = fmt.Sprintf("http://%s",domain)
	}
	u,err := tld.Parse(domain)
	if err!=nil{
		return nil,err
	}
	return u,nil
}

func (f *Factory)CreateBaiDuSpider(domain string) crawler.SpiderInterface{
	u,err := f.validDomain(domain)
	if err != nil{
		panic("初始化爬虫失败")
	}
	return &BaiDuSpider{
		baseSpider:baseSpider{
			domain:fmt.Sprintf("%s.%s",u.Domain,u.TLD),
			baseUrl: "https://www.baidu.com/s"},
	    domains: map[string]int{}}
}