package crawler

import (
	"fmt"
	"time"
)

type scheduler struct {
	spider    SpiderInterface
	goCnt     int
	taskQueue chan *Task
	sfetch    *fetch
	idle      int
}

type schedulerTask struct {
	t        *Task
	callback chan downResult
}

func NewScheduler(spider SpiderInterface, goCnt int) *scheduler {
	return &scheduler{spider, goCnt, make(chan *Task, goCnt), newFetch(), goCnt}
}

func (s *scheduler) addTask(t *Task) {
	go func(item *Task) {
		// 如果任务的callback为nil，将callback置为爬虫Parse方法
		if item.callback == nil {
			item.callback = s.spider.Parse
		}
		s.taskQueue <- item
	}(t)
}

func (s *scheduler) engine() {
	for i := 0; i < s.goCnt; i++ {
		go func() {
			emptyQueueCnt := 0
		closeLabel:
			for {
				select {
				case t1, isClose := <-s.taskQueue:
					if !isClose {
						fmt.Println("协程退出")
						break closeLabel
					}
					resultChan := make(chan downResult)
					go s.sfetch.down(&schedulerTask{t1, resultChan})
					result := <-resultChan
					close(resultChan)
					cResponse := NewCResponse(result.resp, result.err)
					r := t1.callback(t1, cResponse)
					cResponse.Close()
					if r.TaskData != nil {
						s.addTasks(r.TaskData)
					}
					if r.SetData != nil {
						s.spider.ResultProcess(r.SetData)
					}
				default:
					// 结束协程，当channel持续1分钟未提供消息，我们将退出当前协程
					time.Sleep(10 * time.Second)
					emptyQueueCnt++
					if emptyQueueCnt >= 6 {
						s.Close()
					}
				}
			}
		}()
	}
}

func (s *scheduler) addTasks(tasks []*Task) {
	for _, v := range tasks {
		s.addTask(v)
	}
}

func (s *scheduler) Start() {
	s.addTasks(s.spider.Seeds())
	s.engine()
}

func (s *scheduler) Close() {
	fmt.Println("关闭爬虫")
	close(s.taskQueue)
}
