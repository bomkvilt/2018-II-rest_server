package models

import (
	"strconv"

	"github.com/valyala/fasthttp"
)

// easyjson:json
type Forum struct {
	ID      int    `json:"-"`
	Posts   int    `json:"posts"`
	Slug    string `json:"slug"`
	Threads int    `json:"threads"`
	Title   string `json:"title"`
	Author  string `json:"user"`
}

// easyjson:json
type Forums []*Forum

// easyjson:json
type ForumQuery struct {
	Desc  *bool
	Limit *int
	Since string
	Slug  string
}

func (Forum) FromRequest(ctx *fasthttp.RequestCtx) *Forum {
	u := &Forum{}
	check(u.UnmarshalJSON(ctx.PostBody()))
	return u
}

func (ForumQuery) FromRequest(ctx *fasthttp.RequestCtx) *ForumQuery {
	var (
		limit = string(ctx.QueryArgs().Peek("limit"))
		since = string(ctx.QueryArgs().Peek("since"))
		desc  = string(ctx.QueryArgs().Peek("desc"))
		q     = &ForumQuery{}
	)
	if limit != "" {
		q.Limit = new(int)
		*q.Limit, _ = strconv.Atoi(limit)
	}
	if desc != "" {
		q.Desc = new(bool)
		if desc == "true" {
			*q.Desc = true
		}
	}
	q.Slug = ctx.UserValue("slug").(string)
	q.Since = since
	return q
}
