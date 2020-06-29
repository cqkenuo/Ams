package Crawler

import "net/http"

type spiderInterface interface {
	seeds() []task
	parse(request http.Request,response CResponse)
}

type Spider struct {
	spiderInterface
}

func (s *Spider)start() []task{
	return s.seeds()
}