package models

import (
	"time"

	"github.com/valyala/fasthttp"
)

// easyjson:json
type Thread struct {
	ID      int       `json:"id"`
	Author  string    `json:"author"`
	Created time.Time `json:"created"`
	Forum   string    `json:"forum"`
	Message string    `json:"message"`
	Slug    string    `json:"slug"`
	Title   string    `json:"title"`
	Votes   int       `json:"votes"`
}

// easyjson:json
type Threads []*Thread

func (Thread) FromRequest(ctx *fasthttp.RequestCtx) *Thread {
	u := &Thread{}
	check(u.UnmarshalJSON(ctx.PostBody()))
	return u
}
