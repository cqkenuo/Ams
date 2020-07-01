package crawler

type SpiderInterface interface {
	Seeds() []*Task
	Parse(request *Task, response *CResponse) SpiderResult
	ResultProcess(result []map[string]interface{})
}

type Spider struct {

}