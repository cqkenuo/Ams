package crawler

type SpiderResult struct {
	ResultType int
	TaskData []*Task
	SetData []map[string]interface{}
}