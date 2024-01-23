package requests

import (
	"errors"
	"net/http"
)

type PostStoreRequest struct {
	ThreadID uint   `json:"thread_id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
}

func (p *PostStoreRequest) Bind(r *http.Request) error {
	if p == nil {
		return errors.New("missing required post fields.")
	}
	return nil
}
