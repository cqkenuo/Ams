package domain_info

import (
	"Ams/config"
	"Ams/crawler"
	"Ams/model"
	"encoding/json"
)

func DomainInfoCollect(domain *model.Domains) {
	go func(d *model.Domains) {
		if d.DNS == "[]" || d.DNS == "{}" {
			client := NewDnsResolve()
			resolves := client.LookUp(d.Domain)
			if len(resolves) < 1 {
				resolves = client.LookUp(d.Domain)
			}
			bys, err := json.Marshal(resolves)
			if err == nil {
				db := model.GetAppDB(config.LoadConfig())
				db.Model(&d).Update("DNS", string(bys))
			}
		}
		if d.Title == "[]" || d.Title == "{}" {
			schedule := crawler.NewScheduler(&IndexSpider{Domain: *domain}, 5)
			schedule.Start()
		}
	}(domain)
}
