package Crawler

import "net/http"

type task struct {
	request *http.Request
	callback func(r *http.Request,response *CResponse)
}

func NewTask(r http.Request,callback func(r *http.Request,response *CResponse)) *task {
	return &task{&r,callback}
}
