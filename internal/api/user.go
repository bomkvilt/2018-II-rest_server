package api

import (
	"github.com/valyala/fasthttp"

	"AForum/internal/models"
)

func (h *Handler) UserCreate(ctx *fasthttp.RequestCtx) {
	u := models.User{}.FromRequest(ctx)
	u.Nickname = ctx.UserValue("nick").(string)

	if err := h.db.InsertNewUser(u); err != nil {
		us, _ := h.db.GetAllCollisions(u)
		// println(err.Error(), "ThreadGetOne")
		response(ctx, 409, us)
	} else {
		response(ctx, 201, u)
	}
}

func (h *Handler) UserGetOne(ctx *fasthttp.RequestCtx) {
	nick := ctx.UserValue("nick").(string)

	if u, err := h.db.GetUserByName(nick); err != nil {
		// println(err.Error(), "UserGetOne")
		response(ctx, 404, models.Error{Message: "Can't find user by nickname: " + nick})
	} else {
		response(ctx, 200, u)
	}
}

func (h *Handler) UserUpdate(ctx *fasthttp.RequestCtx) {
	u := models.User{}.FromRequest(ctx)
	u.Nickname = ctx.UserValue("nick").(string)

	// no author found
	if _, err := h.db.GetUserByName(u.Nickname); err != nil {
		// println(err.Error(), "UserUpdate")
		response(ctx, 404, models.Error{Message: "Can't find user by nickname: " + u.Nickname})
		return
	}
	// sucess update
	if err := h.db.UpdateUser(u); err == nil {
		u, _ = h.db.GetUserByName(u.Nickname)
		response(ctx, 200, u)
		return
	} else {
		// the email is already in use
		o, _ := h.db.GetUserByEmail(u.Email)
		// println(err.Error(), "UserUpdate")
		response(ctx, 409, models.Error{Message: "This email is already registered by user: " + o.Nickname})
	}
}
