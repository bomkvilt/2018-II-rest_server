package api

import (
	"github.com/valyala/fasthttp"
	"strconv"
	"strings"
	"time"

	"AForum/internal/database"
	"AForum/internal/models"
)

func (h *Handler) PostsCreate(ctx *fasthttp.RequestCtx) {
	var (
		sod = ctx.UserValue("slug_or_id").(string)
		pss = models.Posts{}.FromRequest(ctx)
	)

	if _, err := h.db.GetThreadBySlugOrID(sod); err != nil {
		// println(err.Error(), "PostsCreate")
		response(ctx, 404, models.Error{Message: "Can't find post thread by id: " + sod})
		return
	}
	created := time.Now().Format("2006-01-02T15:04:05.999999999Z07:00")
	
	switch err := h.db.CreateNewPosts(sod, created, pss); err.(type) {
	case *database.ErrorConflict:
		// println(err.Error(), "PostsCreate")
		response(ctx, 409, models.Error{Message: err.Error()})
	case *database.ErrorNotFound:
		// println(err.Error(), "PostsCreate")
		response(ctx, 404, models.Error{Message: err.Error()})
	default:
		response(ctx, 201, pss)
	}
}

func (h *Handler) ThreadGetPosts(ctx *fasthttp.RequestCtx) {
	q := models.PostQuery{}.FromRequest(ctx)
	if pss, err := h.db.GetPosts(q); err != nil {
		// println(err.Error(), "ThreadGetPosts")
		response(ctx, 404, models.Error{Message: err.Error()})
	} else {
		response(ctx, 200, pss)
	}
}

func (h *Handler) PostGetOne(ctx *fasthttp.RequestCtx) {
	var (
		rawID   = ctx.UserValue("id").(string)
		rawRel  = ctx.QueryArgs().Peek("related")
		related = strings.Split(string(rawRel), ",")
		id, _   = strconv.ParseInt(rawID, 10, 64)
	)
	if res, err := h.db.GetPost(id, related); err != nil {
		// println(err.Error(), "PostGetOne")
		response(ctx, 404, models.Error{Message: err.Error()})
	} else {
		response(ctx, 200, res)
	}
}

func (h *Handler) PostUpdate(ctx *fasthttp.RequestCtx) {
	var (
		p = models.Post{}.FromRequest(ctx)
		rawID   = ctx.UserValue("id").(string)
	)
	p.ID, _ = strconv.ParseInt(rawID, 10, 64)

	if err := h.db.UpdatePost(p); err != nil {
		// println(err.Error(), "PostUpdate")
		response(ctx, 404, models.Error{Message: err.Error()})
	} else {
		response(ctx, 200, p)
	}
}
