package api

import (
	"AForum/internal/database"
	"encoding/json"

	"github.com/valyala/fasthttp"
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

func (h *Handler) Clear(ctx *fasthttp.RequestCtx) {
	h.db.Clear()
	response(ctx, 200, nil)
}

func (h *Handler) Status(ctx *fasthttp.RequestCtx) {
	response(ctx, 200, h.db.GetStatus())
}

// ----------------|

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func response(ctx *fasthttp.RequestCtx, code int, payload json.Marshaler) {
	if payload != nil {
		b, _ := payload.MarshalJSON()
		ctx.SetContentType("application/json")
		ctx.SetStatusCode(code)
		ctx.Write(b)
	} else {
		ctx.SetStatusCode(code)
	}
}
