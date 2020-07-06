package crawler

import (
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

type CResponse struct {
	*http.Response
	Extract *goquery.Document
	HttpErr error
	Err     error
}

func NewCResponse(response *http.Response, httpErr error) *CResponse {
	if httpErr != nil {
		return &CResponse{Response: response, HttpErr: httpErr}
	}
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return &CResponse{Response: response, Err: err, HttpErr: httpErr}
	}
	return &CResponse{Response: response, Extract: doc, HttpErr: httpErr}
}

func (cr *CResponse) Close() {
	cr.Body.Close()
}
