package api

import (
	"net/http"

	"github.com/gorilla/mux"

	"AForum/internal/database"
	"AForum/internal/models"
)

func (h *Handler) ForumCreate(rw http.ResponseWriter, r *http.Request) {
	f := models.Forum{}.FromRequest(r)

	switch err := h.db.CreateNewForum(f); err.(type) {
	case *database.ErrorNotFound:
		println(err.Error(), "ForumCreate")
		response(rw, 404, models.Error{Message: "Can't find user with nickname: " + f.Author})
	case *database.ErrorAlreadyExist:
		println(err.Error(), "ForumCreate")
		response(rw, 409, f)
	default:
		response(rw, 201, f)
	}
}

func (h *Handler) ForumGetOne(rw http.ResponseWriter, r *http.Request) {
	if f, err := h.db.GetForumBySlug(mux.Vars(r)["slug"]); err == nil {
		response(rw, 200, f)
	} else {
		println(err.Error(), "ForumGetOne")
		response(rw, 404, models.Error{Message: err.Error()})
	}
}

func (h *Handler) ForumGetUsers(rw http.ResponseWriter, r *http.Request) {
	q := models.ForumQuery{}.FromRequest(r)
	res, err := h.db.GetForumUsers(q)
	if err != nil {
		println(err.Error(), "ForumGetUsers")
		response(rw, 404, models.Error{Message: err.Error()})
		return
	}
	if res == nil {
		res = models.Users{}
	}
	response(rw, 200, res)
}
