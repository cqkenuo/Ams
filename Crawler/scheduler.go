package Crawler

import "container/list"
type scheduler struct {
	spiders chan Spider
	taskQueue *list.List
}

func NewScheduler(spiderChan chan Spider) *scheduler {
	return &scheduler{spiderChan,list.New()}
}

func (s *scheduler)Start(){
	for{
		spider := <- s.spiders
		for _,v := range spider.start(){
			s.taskQueue.PushBack(v)
		}
	}
}