package api

import (
	"AForum/internal/database"
	"encoding/json"
	"net/http"
)

type Handler struct {
	db *database.DB
}

func New(db *database.DB) *Handler {
	return &Handler{
		db: db,
	}
}

// ----------------|

func (h *Handler) Clear(rw http.ResponseWriter, r *http.Request) {
	h.db.Clear()
	response(rw, 200, nil)
}

func (h *Handler) Status(rw http.ResponseWriter, r *http.Request) {
	response(rw, 200, h.db.GetStatus())
}

// ----------------|

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func response(rw http.ResponseWriter, code int, payload interface{}) {
	rw.WriteHeader(code)
	if payload == nil {
		return
	}
	b, err := json.Marshal(payload)
	check(err)
	_, err = rw.Write(b)
	check(err)
}
