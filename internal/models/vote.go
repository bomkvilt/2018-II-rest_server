package models

import (
	"io/ioutil"
	"net/http"
)

// easyjson:json
type Vote struct {
	ID       int    `json:"id"`
	Nickname string `json:"nickname"`
	Voice    int    `json:"voice"`
}

func (Vote) FromRequest(r *http.Request) *Vote {
	b, err := ioutil.ReadAll(r.Body)
	check(err)

	u := &Vote{}
	check(u.UnmarshalJSON(b))
	return u
}
