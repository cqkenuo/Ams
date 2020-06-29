package Crawler
import (
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

type CResponse struct {
	*http.Response
	extract *goquery.Document
}

func NewCResponse(response *http.Response) *CResponse {
	doc,err := goquery.NewDocumentFromReader(response.Body)
	if err != nil{
		panic("响应错误")
	}
	return &CResponse{response,doc}
}

func (cr *CResponse)Close() {
	cr.Body.Close()
}