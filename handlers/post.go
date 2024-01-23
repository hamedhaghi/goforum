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

func PostList(ps services.PostServiceInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pp, err := ps.FindAll()
		if err != nil {
			toJson(w, r, err, http.StatusInternalServerError)
			return
		}
		res := responses.PostListResponse(*pp)
		toJson(w, r, res, http.StatusOK)
	}
}

func PostShow(ps services.PostServiceInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			toJson(w, r, nil, http.StatusBadRequest)
			return
		}
		p, err := ps.FindByID(uint(id))
		if err != nil {
			toJson(w, r, err, http.StatusInternalServerError)
			return
		}
		res := responses.PostSingleResponse(*p)
		toJson(w, r, res, http.StatusOK)
	}
}

func PostStore(ps services.PostServiceInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &requests.PostStoreRequest{}
		if err := render.Bind(r, req); err != nil {
			toJson(w, r, nil, http.StatusBadRequest)
			return
		}
		if req.ThreadID == 0 {
			toJson(w, r, nil, http.StatusBadRequest)
			return
		}
		err := ps.Create(&models.Post{
			ThreadID: req.ThreadID,
			Title:    req.Title,
			Content:  req.Content,
		})
		if err != nil {
			toJson(w, r, err, http.StatusInternalServerError)
			return
		}
		toJson(w, r, nil, http.StatusCreated)
	}
}

func PostUpdate(ps services.PostServiceInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			toJson(w, r, nil, http.StatusBadRequest)
			return
		}
		req := &requests.PostStoreRequest{}
		if err := render.Bind(r, req); err != nil {
			toJson(w, r, nil, http.StatusBadRequest)
			return
		}
		post, err := ps.FindByID(uint(id))
		if err != nil {
			toJson(w, r, err, http.StatusInternalServerError)
			return
		}
		post.ThreadID = req.ThreadID
		post.Title = req.Title
		post.Content = req.Content
		err = ps.Update(post)
		if err != nil {
			toJson(w, r, err, http.StatusInternalServerError)
			return
		}
		toJson(w, r, req, http.StatusOK)
	}
}

func PostDelete(ps services.PostServiceInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			toJson(w, r, nil, http.StatusBadRequest)
			return
		}
		err = ps.Delete(uint(id))
		if err != nil {
			toJson(w, r, err, http.StatusInternalServerError)
			return
		}
		toJson(w, r, nil, http.StatusNoContent)
	}
}
