package Crawler

type SpiderResult struct {
	ResultType int
	TaskData []*Task
	SetData []map[string]interface{}
}