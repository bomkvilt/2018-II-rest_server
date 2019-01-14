package models

import "github.com/valyala/fasthttp"

// easyjson:json
type User struct {
	ID       int64  `json:"-"`
	About    string `json:"about"`
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
	Nickname string `json:"nickname"`
}

// easyjson:json
type Users []*User

func (User) FromRequest(ctx *fasthttp.RequestCtx) *User {
	u := &User{}
	check(u.UnmarshalJSON(ctx.PostBody()))
	return u
}
