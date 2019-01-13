package api

import (
	"net/http"

	"github.com/gorilla/mux"

	"AForum/internal/database"
	"AForum/internal/models"
)

func (h *Handler) ForumCreate(rw http.ResponseWriter, r *http.Request) {
	f := models.Forum{}.FromRequest(r)

	switch h.db.CreateNewForum(f).(type) {
	case *database.ErrorNotFound:
		response(rw, 404, models.Error{Message: "Can't find user with nickname: " + f.Author})
	case *database.ErrorAlreadyExist:
		response(rw, 409, f)
	default:
		response(rw, 201, f)
	}
}

func (h *Handler) ForumGetOne(rw http.ResponseWriter, r *http.Request) {
	if f, err := h.db.GetForumBySlug(mux.Vars(r)["slug"]); err == nil {
		response(rw, 200, f)
		return
	}
	response(rw, 404, models.Error{Message: ""})
}

func (h *Handler) ForumGetUsers(rw http.ResponseWriter, r *http.Request) {
	q := models.ForumQuery{}.FromRequest(r)
	res, err := h.db.GetForumUsers(q)
	if err != nil {
		response(rw, 404, models.Error{Message: ""})
		return
	}
	response(rw, 200, res)
}
