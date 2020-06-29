package Crawler

type scheduler struct {
	spider    *Spider
	goCnt     int
	taskQueue chan *task
	sfetch    *fetch
}

type schedulerTask struct {
	t        *task
	callback chan *downResult
}

func NewScheduler(spider *Spider, goCnt int) *scheduler {
	return &scheduler{spider, goCnt, make(chan *task, goCnt), newFetch()}
}

func (s *scheduler) addTask(t *task) {
	go func(item *task) {
		s.taskQueue <- item
	}(t)
}

func (s *scheduler) engine() {
	for i := 0; i < s.goCnt; i++ {
		go func() {
			for {
				t := <-s.taskQueue
				resultChan := make(chan *downResult)
				s.sfetch.down(&schedulerTask{t, resultChan})
				result := <-resultChan
				t.callback(t.request, NewCResponse(result.resp))
			}
		}()
	}
}

func (s *scheduler) Start() {
	for _, v := range s.spider.start() {
		s.addTask(v)
	}
	s.engine()
}
