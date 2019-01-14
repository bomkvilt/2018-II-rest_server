package api

import (
	"github.com/valyala/fasthttp"

	"AForum/internal/database"
	"AForum/internal/models"
)

func (h *Handler) ForumCreate(ctx *fasthttp.RequestCtx) {
	f := models.Forum{}.FromRequest(ctx)

	switch err := h.db.CreateNewForum(f); err.(type) {
	case *database.ErrorNotFound:
		// println(err.Error(), "ForumCreate")
		response(ctx, 404, models.Error{Message: "Can't find user with nickname: " + f.Author})
	case *database.ErrorAlreadyExist:
		// println(err.Error(), "ForumCreate")
		response(ctx, 409, f)
	default:
		response(ctx, 201, f)
	}
}

func (h *Handler) ForumGetOne(ctx *fasthttp.RequestCtx) {
	slug := ctx.UserValue("slug").(string)

	if f, err := h.db.GetForumBySlug(slug); err == nil {
		response(ctx, 200, f)
	} else {
		// println(err.Error(), "ForumGetOne")
		response(ctx, 404, models.Error{Message: err.Error()})
	}
}

func (h *Handler) ForumGetUsers(ctx *fasthttp.RequestCtx) {
	q := models.ForumQuery{}.FromRequest(ctx)
	res, err := h.db.GetForumUsers(q)
	if err != nil {
		// println(err.Error(), "ForumGetUsers")
		response(ctx, 404, models.Error{Message: err.Error()})
		return
	}
	if res == nil {
		res = models.Users{}
	}
	response(ctx, 200, res)
}
