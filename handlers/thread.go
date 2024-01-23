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

func ThreadList(ts services.ThreadServiceInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		threads, err := ts.FindAll()
		if err != nil {
			toJson(w, r, err.Error(), http.StatusInternalServerError)
			return
		}
		res := responses.ThreadListResponse(*threads)
		toJson(w, r, res, http.StatusOK)
	}
}

func ThreadShow(ts services.ThreadServiceInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			toJson(w, r, nil, http.StatusBadRequest)
			return
		}
		t, err := ts.FindByID(uint(id))
		if err != nil {
			toJson(w, r, err, http.StatusInternalServerError)
			return
		}
		res := responses.ThreadSingleResponse(*t)
		toJson(w, r, res, http.StatusOK)
	}
}

func ThreadStore(ts services.ThreadServiceInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &requests.ThreadStoreRequest{}
		if err := render.Bind(r, req); err != nil {
			toJson(w, r, nil, http.StatusBadRequest)
			return
		}
		err := ts.Create(&models.Thread{
			Title:       req.Title,
			Description: req.Description,
		})
		if err != nil {
			toJson(w, r, err, http.StatusInternalServerError)
			return
		}
		toJson(w, r, nil, http.StatusCreated)
	}
}

func ThreadUpdate(ts services.ThreadServiceInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			toJson(w, r, err, http.StatusBadRequest)
			return
		}
		req := &requests.ThreadStoreRequest{}
		if err := render.Bind(r, req); err != nil {
			toJson(w, r, nil, http.StatusBadRequest)
			return
		}
		thread, err := ts.FindByID(uint(id))
		if err != nil {
			toJson(w, r, err, http.StatusInternalServerError)
			return
		}
		thread.Title = req.Title
		thread.Description = req.Description
		err = ts.Update(thread)
		if err != nil {
			toJson(w, r, err, http.StatusInternalServerError)
			return
		}
		toJson(w, r, req, http.StatusOK)
	}
}

func ThreadDelete(ts services.ThreadServiceInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			toJson(w, r, err, http.StatusBadRequest)
			return
		}
		err = ts.Delete(uint(id))
		if err != nil {
			toJson(w, r, err, http.StatusInternalServerError)
			return
		}
		toJson(w, r, nil, http.StatusNoContent)
	}
}
