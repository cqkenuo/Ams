package Crawler

import "net/http"

type Task struct {
	request *http.Request
	callback func(r *http.Request,response *CResponse) SpiderResult
}

func NewTask(r *http.Request,callback func(r *http.Request,response *CResponse) SpiderResult) *Task {
	return &Task{r,callback}
}