package domain_info

import (
	"Ams/config"
	"Ams/crawler"
	"Ams/model"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strconv"
	"sync"
)

type IndexSpider struct {
	Domain model.Domains
	result map[string]map[string]string
	lock   sync.Mutex
}

func (i *IndexSpider) Seeds() []*crawler.Task {
	// 初始化result
	i.result = make(map[string]map[string]string)
	tasks := make([]*crawler.Task, 2)
	r, _ := http.NewRequest("GET", fmt.Sprintf("http://%s", i.Domain.Domain), nil)
	tasks[0] = crawler.NewTask(r, nil)
	tasks[0].SetAttach(map[string]interface{}{"protocol": "http"})
	r1, _ := http.NewRequest("GET", fmt.Sprintf("https://%s", i.Domain.Domain), nil)
	tasks[1] = crawler.NewTask(r1, nil)
	tasks[1].SetAttach(map[string]interface{}{"protocol": "https"})
	return tasks
}

func (i IndexSpider) Parse(request *crawler.Task, response *crawler.CResponse) crawler.SpiderResult {
	//tasks := make([]*crawler.Task, 0)
	sets := make([]map[string]interface{}, 0)
	if response.Err != nil {
		sets = append(sets, map[string]interface{}{"err": response.Err.Error(), "protocol": request.GetAttach()["protocol"]})
		return crawler.SpiderResult{SetData: sets}
	}
	if response.HttpErr != nil {
		sets = append(sets, map[string]interface{}{"err": response.HttpErr.Error(), "protocol": request.GetAttach()["protocol"]})
		return crawler.SpiderResult{SetData: sets}
	}
	result := make(map[string]interface{})
	response.Extract.Find("title").Each(func(i int, selection *goquery.Selection) {
		result = map[string]interface{}{"title": selection.Text()}
	})
	result["status"] = strconv.FormatInt(int64(response.StatusCode), 10)
	result["protocol"] = request.GetAttach()["protocol"]
	sets = append(sets, result)
	return crawler.SpiderResult{SetData: sets}
}

func (i *IndexSpider) ResultProcess(result []map[string]interface{}) {
	for _, dict := range result {
		if dict == nil {
			continue
		}
		err, errOk := dict["err"]
		protocol := dict["protocol"]
		if protocol == nil {
			continue
		}
		tmpMap := make(map[string]string)
		if errOk {
			tmpMap["error"] = err.(string)
		}
		title, titleOk := dict["title"]
		if titleOk {
			tmpMap["title"] = title.(string)
		}
		status, statusOk := dict["status"]
		if statusOk {
			tmpMap["status"] = status.(string)
		}
		i.lock.Lock()
		i.result[protocol.(string)] = tmpMap
		i.lock.Unlock()
	}
}

func (i *IndexSpider) Close() {
	fmt.Println("首页爬虫结束:", i.Domain.Domain)
	bys, err := json.Marshal(i.result)
	if err == nil {
		db := model.GetAppDB(config.LoadConfig())
		db.Model(&i.Domain).Update("Title", string(bys))
	}
}
