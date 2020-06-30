package Crawler

import (
	"net/http"
)

type SpiderInterface interface {
	Seeds() []*Task
	Parse(request *http.Request, response *CResponse) SpiderResult
	ResultProcess(result []map[string]interface{})
}

type Spider struct {

}