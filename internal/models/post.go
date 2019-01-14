package models

import (
	"strconv"
	"time"

	"github.com/valyala/fasthttp"
)

// easyjson:json
type Post struct {
	ID       int64     `json:"id,omitempty"`
	Author   string    `json:"author"`
	Created  time.Time `json:"created"`
	Forum    string    `json:"forum,omitempty"`
	IsEdited bool      `json:"isEdited,omitempty"`
	Message  string    `json:"message"`
	Parent   int64     `json:"parent,omitempty"`
	Thread   int       `json:"thread,omitempty"`
	Path     []int64   `json:"-"`
}

// easyjson:json
type Posts []*Post

// easyjson:json
type PostFull struct {
	Author *User   `json:"author,omitempty"`
	Forum  *Forum  `json:"forum,omitempty"`
	Post   *Post   `json:"post,omitempty"`
	Thread *Thread `json:"thread,omitempty"`
}

// easyjson:json
type PostQuery struct {
	Desc     *bool
	Limit    *int
	Since    int64
	SlugOrID string
	Sort     string
}

func (Post) FromRequest(ctx *fasthttp.RequestCtx) *Post {
	u := &Post{}
	check(u.UnmarshalJSON(ctx.PostBody()))
	return u
}

func (Posts) FromRequest(ctx *fasthttp.RequestCtx) Posts {
	u := Posts{}
	check(u.UnmarshalJSON(ctx.PostBody()))
	return u
}

func (PostQuery) FromRequest(ctx *fasthttp.RequestCtx) *PostQuery {
	var (
		q = &PostQuery{}
	)
	if ctx.QueryArgs().Has("limit") {
		q.Limit = new(int)
		*q.Limit = ctx.QueryArgs().GetUintOrZero("limit")
	}
	if ctx.QueryArgs().Has("since") {
		raw := string(ctx.QueryArgs().Peek("since"))
		q.Since, _ = strconv.ParseInt(raw, 10, 64)
	}
	if ctx.QueryArgs().Has("desc") {
		q.Desc = new(bool)
		*q.Desc = ctx.QueryArgs().GetBool("desc")
	}
	if ctx.QueryArgs().Has("sort") {
		q.Sort = string(ctx.QueryArgs().Peek("sort"))
	} else {
		q.Sort = "flat"
	}
	q.SlugOrID = ctx.UserValue("slug_or_id").(string)
	return q
}
