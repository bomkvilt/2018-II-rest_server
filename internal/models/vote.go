package models

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Vote struct {
	ID       int    `json:"id"`
	Nickname string `json:"nickname"`
	Voice    int    `json:"voice"`
}

func (Vote) FromRequest(r *http.Request) *Vote {
	b, err := ioutil.ReadAll(r.Body)
	check(err)

	u := &Vote{}
	check(json.Unmarshal(b, u))
	return u
}
