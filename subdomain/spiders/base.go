package subdomain

import (
	"Ams/crawler"
	"github.com/jpillora/go-tld"
)

type baseSpider struct {
	crawler.Spider
	baseUrl string
	domain string
}

func (b *baseSpider)TldParse(url string)*tld.URL{
	u,_ := tld.Parse(url)
	return u
}


