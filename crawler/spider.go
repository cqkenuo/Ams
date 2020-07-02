package crawler

type SpiderInterface interface {
	Seeds() []*Task
	Parse(request *Task, response *CResponse) SpiderResult
	ResultProcess(result []map[string]interface{})
	Close()
}

type Spider struct {

}