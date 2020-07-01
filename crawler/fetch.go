package crawler

import (
	"net/http"
)

type fetch struct {
	client *http.Client
}

func newFetch() *fetch {
	return &fetch{&http.Client{}}
}

type downResult struct {
	resp *http.Response
	err  error
}

func (f *fetch) down(t *schedulerTask) {
	resp,err := f.client.Do(t.t.request)
	t.callback <- downResult{resp,err}
}
