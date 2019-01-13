package models

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

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

type Threads []*Thread

func (Thread) FromRequest(r *http.Request) *Thread {
	b, err := ioutil.ReadAll(r.Body)
	check(err)

	u := &Thread{}
	check(json.Unmarshal(b, u))
	return u
}
