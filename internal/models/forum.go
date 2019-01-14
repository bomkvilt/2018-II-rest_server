package models

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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
	Since *string
	Slug  string
}

func (Forum) FromRequest(r *http.Request) *Forum {
	b, err := ioutil.ReadAll(r.Body)
	check(err)

	u := &Forum{}
	check(json.Unmarshal(b, u))
	return u
}

func (ForumQuery) FromRequest(r *http.Request) *ForumQuery {
	var (
		limit = r.URL.Query().Get("limit")
		since = r.URL.Query().Get("since")
		desc  = r.URL.Query().Get("desc")
		q     = &ForumQuery{}
	)
	q.Slug = mux.Vars(r)["slug"]
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
	if since != "" {
		q.Since = new(string)
		*q.Since = since
	}
	return q
}
