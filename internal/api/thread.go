package api

import (
	"net/http"

	"github.com/gorilla/mux"

	"AForum/internal/database"
	"AForum/internal/models"
)

func (h *Handler) ThreadCreate(rw http.ResponseWriter, r *http.Request) {
	var (
		th   = models.Thread{}.FromRequest(r)
		slug = mux.Vars(r)["slug"]
	)
	switch err := h.db.CreateNewThread(slug, th); err.(type) {
	case *database.ErrorNotFound:
		response(rw, 404, models.Error{Message: err.Error()})
	case *database.ErrorAlreadyExist:
		response(rw, 409, th)
	default:
		response(rw, 201, th)
	}
}

func (h *Handler) ThreadUpdate(rw http.ResponseWriter, r *http.Request) {
	var (
		th  = models.Thread{}.FromRequest(r)
		sod = mux.Vars(r)["slug_or_id"]
	)
	if err := h.db.UpdateThread(sod, th); err != nil {
		response(rw, 404, models.Error{Message: err.Error()})
	} else {
		response(rw, 200, th)
	}
}

func (h *Handler) ThreadGetOne(rw http.ResponseWriter, r *http.Request) {
	th, err := h.db.GetThreadBySlugOrID(mux.Vars(r)["slug_or_id"])
	if err != nil {
		response(rw, 404, models.Error{Message: err.Error()})
	} else {
		response(rw, 200, th)
	}
}

func (h *Handler) ForumGetThreads(rw http.ResponseWriter, r *http.Request) {
	q := models.ForumQuery{}.FromRequest(r)
	if ths, err := h.db.GetThreads(q); err != nil {
		response(rw, 404, models.Error{Message: err.Error()})
	} else {
		response(rw, 200, ths)
	}
}

func (h *Handler) ThreadVote(rw http.ResponseWriter, r *http.Request) {
	var (
		vt  = models.Vote{}.FromRequest(r)
		sod = mux.Vars(r)["slug_or_id"]
	)
	if th, err := h.db.VoteThread(sod, vt); err != nil {
		response(rw, 404, models.Error{Message: err.Error()})
	} else {
		response(rw, 200, th)
	}
}
