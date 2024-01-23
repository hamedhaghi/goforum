package handlers

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"

	"github.com/hamedhaghi/goforum/handlers/requests"
	"github.com/hamedhaghi/goforum/handlers/responses"
	"github.com/hamedhaghi/goforum/models"
	"github.com/hamedhaghi/goforum/services"
)

func CommentList(cs services.CommentServiceInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cc, err := cs.FindAll()
		if err != nil {
			toJson(w, r, err, http.StatusInternalServerError)
			return
		}
		res := responses.CommentListResponse(*cc)
		toJson(w, r, res, http.StatusOK)
	}
}

func CommentShow(cs services.CommentServiceInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			toJson(w, r, nil, http.StatusBadRequest)
			return
		}
		c, err := cs.FindByID(uint(id))
		if err != nil {
			toJson(w, r, err, http.StatusInternalServerError)
			return
		}
		res := responses.CommentSingleResponse(*c)
		toJson(w, r, res, http.StatusOK)
	}
}

func CommentStore(cs services.CommentServiceInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &requests.CommentStoreRequest{}
		if err := render.Bind(r, req); err != nil {
			toJson(w, r, nil, http.StatusBadRequest)
			return
		}
		err := cs.Create(&models.Comment{
			PostID:  req.PostID,
			Content: req.Content,
		})
		if err != nil {
			toJson(w, r, err, http.StatusInternalServerError)
			return
		}
		toJson(w, r, nil, http.StatusCreated)
	}
}

func CommentUpdate(cs services.CommentServiceInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			toJson(w, r, nil, http.StatusBadRequest)
			return
		}
		req := &requests.CommentStoreRequest{}
		if err := render.Bind(r, req); err != nil {
			toJson(w, r, nil, http.StatusBadRequest)
			return
		}
		comment, err := cs.FindByID(uint(id))
		if err != nil {
			toJson(w, r, err, http.StatusInternalServerError)
			return
		}
		comment.PostID = req.PostID
		comment.Content = req.Content
		err = cs.Update(comment)
		if err != nil {
			toJson(w, r, err, http.StatusInternalServerError)
			return
		}
		toJson(w, r, req, http.StatusOK)
	}
}

func CommentDelete(cs services.CommentServiceInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			toJson(w, r, nil, http.StatusBadRequest)
			return
		}
		err = cs.Delete(uint(id))
		if err != nil {
			toJson(w, r, err, http.StatusInternalServerError)
			return
		}
		toJson(w, r, nil, http.StatusNoContent)
	}
}
