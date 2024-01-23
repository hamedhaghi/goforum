package requests

import (
	"errors"
	"net/http"
)

type CommentStoreRequest struct {
	PostID  uint   `json:"post_id"`
	Content string `json:"content"`
}

func (c *CommentStoreRequest) Bind(r *http.Request) error {
	if c == nil {
		return errors.New("missing required comment fields.")
	}
	return nil
}
