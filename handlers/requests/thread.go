package requests

import (
	"errors"
	"net/http"
)

type ThreadStoreRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (t *ThreadStoreRequest) Bind(r *http.Request) error {
	if t == nil {
		return errors.New("missing required thread fields.")
	}
	return nil
}
