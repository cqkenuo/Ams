package crawler

import (
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

type CResponse struct {
	*http.Response
	Extract *goquery.Document
	httpErr error
	err     error
}

func NewCResponse(response *http.Response, httpErr error) *CResponse {
	if httpErr != nil {
		return &CResponse{Response: response, httpErr: httpErr}
	}
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return &CResponse{Response: response, err: err, httpErr: httpErr}
	}
	return &CResponse{Response: response, Extract: doc, httpErr: httpErr}
}

func (cr *CResponse) Close() {
	cr.Body.Close()
}
