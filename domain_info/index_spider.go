package domain_info

import (
	"Ams/crawler"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strconv"
)

type IndexSpider struct {
	Domain string
}

func (i IndexSpider) Seeds() []*crawler.Task {
	tasks := make([]*crawler.Task, 2)
	r, _ := http.NewRequest("GET", fmt.Sprintf("http://%s", i.Domain), nil)
	tasks[0] = crawler.NewTask(r, nil)
	r1, _ := http.NewRequest("GET", fmt.Sprintf("https://%s", i.Domain), nil)
	tasks[1] = crawler.NewTask(r1, nil)
	return tasks
}

func (i IndexSpider) Parse(request *crawler.Task, response *crawler.CResponse) crawler.SpiderResult {
	//tasks := make([]*crawler.Task, 0)
	sets := make([]map[string]interface{}, 0)
	fmt.Println(response.Err)
	if response.Err != nil {
		sets = append(sets, map[string]interface{}{"err": response.Err.Error()})
		return crawler.SpiderResult{SetData: sets}
	}
	if response.HttpErr != nil {
		sets = append(sets, map[string]interface{}{"err": response.HttpErr.Error()})
		return crawler.SpiderResult{SetData: sets}
	}
	result := make(map[string]interface{})
	response.Extract.Find("title").Each(func(i int, selection *goquery.Selection) {
		result = map[string]interface{}{"status": strconv.FormatInt(int64(response.StatusCode), 10), "title": selection.Text()}
	})
	sets = append(sets, result)
	return crawler.SpiderResult{SetData: sets}
}

func (i IndexSpider) ResultProcess(result []map[string]interface{}) {
	for k, v := range result {
		fmt.Println(k, v)
	}
}

func (i IndexSpider) Close() {
	panic("implement me")
}
