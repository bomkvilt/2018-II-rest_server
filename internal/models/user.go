package models

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

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

func (User) FromRequest(r *http.Request) *User {
	b, err := ioutil.ReadAll(r.Body)
	check(err)

	u := &User{}
	check(json.Unmarshal(b, u))
	return u
}
