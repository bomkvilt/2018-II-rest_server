package api

import (
	"net/http"

	"github.com/gorilla/mux"

	"AForum/internal/models"
)

func (h *Handler) UserCreate(rw http.ResponseWriter, r *http.Request) {
	u := models.User{}.FromRequest(r)
	u.Nickname = mux.Vars(r)["nick"]
	if err := h.db.InsertNewUser(u); err != nil {
		us, _ := h.db.GetAllCollisions(u)
		response(rw, 409, us)
	} else {
		response(rw, 201, u)
	}
}

func (h *Handler) UserGetOne(rw http.ResponseWriter, r *http.Request) {
	nick := mux.Vars(r)["nick"]

	if u, err := h.db.GetUserByName(nick); err != nil {
		response(rw, 404, models.Error{Message: "Can't find user by nickname: " + nick})
	} else {
		response(rw, 200, u)
	}
}

func (h *Handler) UserUpdate(rw http.ResponseWriter, r *http.Request) {
	u := models.User{}.FromRequest(r)
	u.Nickname = mux.Vars(r)["nick"]
	// no author found
	if _, err := h.db.GetUserByName(u.Nickname); err != nil {
		response(rw, 404, models.Error{Message: "Can't find user by nickname: " + u.Nickname})
		return
	}
	// sucess update
	if err := h.db.UpdateUser(u); err == nil {
		u, _ = h.db.GetUserByName(u.Nickname)
		response(rw, 200, u)
		return
	}
	// the email is already in use
	o, _ := h.db.GetUserByEmail(u.Email)
	response(rw, 409, models.Error{Message: "This email is already registered by user: " + o.Nickname})
}
