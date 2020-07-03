package crawler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Task struct {
	request  *http.Request
	callback func(r *Task, response *CResponse) SpiderResult
	attach   []byte
}

func NewTask(r *http.Request, callback func(r *Task, response *CResponse) SpiderResult) *Task {
	return &Task{request: r, callback: callback}
}

func (t *Task) SetAttach(value map[string]interface{}) {
	jsonStr, err := json.Marshal(value)
	if err != nil {
		fmt.Println("Task Map to Json Error")
	}
	t.attach = jsonStr
}

func (t *Task) GetAttach() map[string]interface{} {
	var result map[string]interface{}
	err := json.Unmarshal(t.attach, &result)
	if err != nil {
		fmt.Println("Task Json to Map Error")
	}
	return result
}
