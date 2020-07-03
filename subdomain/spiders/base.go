package subdomain

import (
	"Ams/crawler"
	"Ams/model"
	"github.com/jpillora/go-tld"
	"net/http"
)

type baseSpider struct {
	crawler.Spider
	baseUrl string
	domain  *model.Domains
}

func (b *baseSpider) TldParse(url string) *tld.URL {
	u, _ := tld.Parse(url)
	return u
}

func (b *baseSpider) NewTask(r *http.Request, c func(task *crawler.Task, response *crawler.CResponse) crawler.SpiderResult) *crawler.Task {
	r.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.106 Safari/537.36")
	r.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	return crawler.NewTask(r, c)
}
