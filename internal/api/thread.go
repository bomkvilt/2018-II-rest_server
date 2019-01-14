package api

import (
	"github.com/valyala/fasthttp"

	"AForum/internal/database"
	"AForum/internal/models"
)

func (h *Handler) ThreadCreate(ctx *fasthttp.RequestCtx) {
	var (
		th   = models.Thread{}.FromRequest(ctx)
		slug = ctx.UserValue("slug").(string)
	)

	switch err := h.db.CreateNewThread(slug, th); err.(type) {
	case *database.ErrorNotFound:
		// println(err.Error(), "ThreadCreate")
		response(ctx, 404, models.Error{Message: err.Error()})
	case *database.ErrorAlreadyExist:
		// println(err.Error(), "ThreadCreate")
		response(ctx, 409, th)
	default:
		response(ctx, 201, th)
	}
}

func (h *Handler) ThreadUpdate(ctx *fasthttp.RequestCtx) {
	var (
		th  = models.Thread{}.FromRequest(ctx)
		sod = ctx.UserValue("slug_or_id").(string)
	)

	if err := h.db.UpdateThread(sod, th); err != nil {
		// println(err.Error(), "ThreadUpdate")
		response(ctx, 404, models.Error{Message: err.Error()})
	} else {
		response(ctx, 200, th)
	}
}

func (h *Handler) ThreadGetOne(ctx *fasthttp.RequestCtx) {
	var (
		sod = ctx.UserValue("slug_or_id").(string)
		th, err = h.db.GetThreadBySlugOrID(sod)
	)

	if err != nil {
		// println(err.Error(), "ThreadGetOne")
		response(ctx, 404, models.Error{Message: err.Error()})
	} else {
		response(ctx, 200, th)
	}
}

func (h *Handler) ForumGetThreads(ctx *fasthttp.RequestCtx) {
	q := models.ForumQuery{}.FromRequest(ctx)

	if ths, err := h.db.GetThreads(q); err != nil {
		// println(err.Error(), "ForumGetThreads")
		response(ctx, 404, models.Error{Message: err.Error()})
	} else {
		response(ctx, 200, ths)
	}
}

func (h *Handler) ThreadVote(ctx *fasthttp.RequestCtx) {
	var (
		vt  = models.Vote{}.FromRequest(ctx)
		sod = ctx.UserValue("slug_or_id").(string)
	)

	if th, err := h.db.VoteThread(sod, vt); err != nil {
		// println(err.Error(), "ThreadVote")
		response(ctx, 404, models.Error{Message: err.Error()})
	} else {
		response(ctx, 200, th)
	}
}
