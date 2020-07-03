package subdomain

import (
	"Ams/crawler"
	"Ams/model"
	"Ams/subdomain/spiders"
	"fmt"
)

type SDServiceTask struct {
	Domain   *model.Domains
	Callback chan []crawler.SpiderInterface
}

func SubdomainService(domainChan chan SDServiceTask) {
	factory := spiders.Factory{}
	for {
		item, ok := <-domainChan
		if ok {
			ss := factory.CreateSpider(item.Domain)
			fmt.Println(ss)
			item.Callback <- ss
		} else {
			fmt.Println("子域名服务驾崩")
			break
		}
	}
}
