package handlers

import (
	"net/http"

	"github.com/go-chi/render"
)

func toJson(w http.ResponseWriter, r *http.Request, data interface{}, statusCode int) {
	render.SetContentType(render.ContentTypeJSON)
	render.Status(r, statusCode)
	render.JSON(w, r, data)
}
