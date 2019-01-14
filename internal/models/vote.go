package models

import "github.com/valyala/fasthttp"

// easyjson:json
type Vote struct {
	ID       int    `json:"id"`
	Nickname string `json:"nickname"`
	Voice    int    `json:"voice"`
}

func (Vote) FromRequest(ctx *fasthttp.RequestCtx) *Vote {
	u := &Vote{}
	check(u.UnmarshalJSON(ctx.PostBody()))
	return u
}
