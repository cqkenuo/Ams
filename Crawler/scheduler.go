package Crawler

type scheduler struct {
	spider    SpiderInterface
	goCnt     int
	taskQueue chan interface{}
	sfetch    *fetch
	idle int
}

type schedulerTask struct {
	t        *Task
	callback chan downResult
}

func NewScheduler(spider SpiderInterface, goCnt int) *scheduler {
	return &scheduler{spider, goCnt, make(chan interface{}, goCnt), newFetch(),goCnt}
}

func (s *scheduler) addTask(t *Task) {
	go func(item *Task) {
		s.taskQueue <- item
	}(t)
}

func (s *scheduler) engine() {
	for i := 0; i < s.goCnt; i++ {
		go func() {
			for {
				t := <-s.taskQueue
				t1,ok := t.(*Task)
				if ok{
					resultChan := make(chan downResult)
					go s.sfetch.down(&schedulerTask{t1, resultChan})
					result := <-resultChan
					r := t1.callback(t1.request, NewCResponse(result.resp))
					if r.ResultType == TaskType {
						s.addTasks(r.TaskData)
					}else if r.ResultType == ResultType {
						s.spider.ResultProcess(r.SetData)
					}
				}else {
					break
				}
			}
		}()
	}
}

func (s *scheduler) addTasks(tasks []*Task){
	for _, v := range tasks {
		s.addTask(v)
	}
}

func (s *scheduler) Start() {
	s.addTasks(s.spider.Seeds())
	s.engine()
}

func (s *scheduler)Close()  {
	for i:=0;i<s.goCnt;i++{
		s.taskQueue <- 1
	}
}