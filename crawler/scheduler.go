package crawler

import (
	"fmt"
	"strings"
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
			emptyQueueCnt := -1
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
					if result.err == nil{
						cResponse.Close()
					}
					if r.TaskData != nil {
						s.addTasks(r.TaskData)
					}
					if r.SetData != nil {
						s.spider.ResultProcess(r.SetData)
					}
				default:
					// 结束协程，当channel持续30秒未提供消息，我们将退出当前协程
					emptyQueueCnt++
					if emptyQueueCnt >= 3 {
						s.Close()
					}
					time.Sleep(10 * time.Second)
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
	defer func() {
		if r := recover(); r != nil {
			// 保证taskQueue只关闭一次 close of closed channel
			if !strings.Contains(r.(error).Error(), "close of closed channel") {
				panic(r)
			}
		}
	}()
	close(s.taskQueue)
	s.spider.Close()
}



func SchedulerService(spiderChan chan SpiderInterface){
	for {
		item,ok := <- spiderChan
		if ok {
			control := NewScheduler(item,10)
			go control.Start()
		}else {
			break
		}
	}
}