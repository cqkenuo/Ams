package subdomain

import (
	"Ams/crawler"
	"Ams/model"
	"fmt"
	"github.com/jpillora/go-tld"
	"reflect"
	"regexp"
	"strings"
)

type Factory struct {
}

func (f *Factory) validDomain(domain string) (*tld.URL, error) {
	if !strings.HasPrefix(domain, "http://") && !strings.HasPrefix(domain, "https://") {
		domain = fmt.Sprintf("http://%s", domain)
	}
	u, err := tld.Parse(domain)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (f *Factory) CreateBaiDuSpider(domain *model.Domains) crawler.SpiderInterface {
	fmt.Println(domain)
	return &BaiDuSpider{
		baseSpider: baseSpider{
			domain:  domain,
			baseUrl: "https://www.baidu.com/s"},
		domains: map[string]int{}}
}

func (f *Factory) CreateSpider(domain *model.Domains) []crawler.SpiderInterface {
	typ := reflect.TypeOf(f)
	val := reflect.ValueOf(f)
	reg, _ := regexp.Compile(`Create[a-zA-Z]+Spider`)
	var spiders []crawler.SpiderInterface
	for i := 0; i < typ.NumMethod(); i++ {
		method := typ.Method(i)
		if reg.Match([]byte(method.Name)) {
			spiders = append(spiders, val.MethodByName(method.Name).Call([]reflect.Value{reflect.ValueOf(domain)})[0].Interface().(crawler.SpiderInterface))
		}
	}
	return spiders
}
