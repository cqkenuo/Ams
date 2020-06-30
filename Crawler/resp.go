package Crawler

import (
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

type CResponse struct {
	*http.Response
	extract *goquery.Document
	err error
}

func NewCResponse(response *http.Response) *CResponse {
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return &CResponse{Response:response,err: err}
	}
	return &CResponse{Response:response,extract: doc}
}

func (cr *CResponse) Close() {
	cr.Body.Close()
}
