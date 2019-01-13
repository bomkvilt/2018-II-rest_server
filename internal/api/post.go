package api

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"

	"AForum/internal/database"
	"AForum/internal/models"
)

func (h *Handler) PostsCreate(rw http.ResponseWriter, r *http.Request) {
	var (
		sod = mux.Vars(r)["slug_or_id"]
		pss = models.Posts{}.FromRequest(r)
	)

	if _, err := h.db.GetThreadBySlugOrID(sod); err != nil {
		response(rw, 404, models.Error{Message: "Can't find post thread by id: " + sod})
	}

	created := time.Now().Format("2006-01-02T15:04:05.999999999Z07:00")
	for _, ps := range pss {
		err := h.db.CreateNewPost(sod, created, ps)
		if err == nil {
			continue
		}

		switch err.(type) {
		case *database.ErrorConflict:
			response(rw, 409, models.Error{Message: err.Error()})
			return
		default:
			response(rw, 404, models.Error{Message: err.Error()})
			return
		}
	}
	response(rw, 201, pss)
}

func (h *Handler) ThreadGetPosts(rw http.ResponseWriter, r *http.Request) {
	q := models.PostQuery{}.FromRequest(r)
	if pss, err := h.db.GetPosts(q); err != nil {
		response(rw, 404, models.Error{Message: ""})
	} else {
		response(rw, 200, pss)
	}
}

func (h *Handler) PostGetOne(rw http.ResponseWriter, r *http.Request) {
	var (
		id, _   = strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
		related = strings.Split(r.URL.Query().Get("related"), ",")
	)
	if res, err := h.db.GetPost(id, related); err != nil {
		response(rw, 404, models.Error{Message: ""})
	} else {
		response(rw, 200, res)
	}
}

func (h *Handler) PostUpdate(rw http.ResponseWriter, r *http.Request) {
	p := models.Post{}.FromRequest(r)
	p.ID, _ = strconv.ParseInt(mux.Vars(r)["id"], 10, 64)

	if err := h.db.UpdatePost(p); err != nil {
		response(rw, 404, models.Error{Message: ""})
	} else {
		response(rw, 200, p)
	}
}
