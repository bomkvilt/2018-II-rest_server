// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// UserGetOneHandlerFunc turns a function with the right signature into a user get one handler
type UserGetOneHandlerFunc func(UserGetOneParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UserGetOneHandlerFunc) Handle(params UserGetOneParams) middleware.Responder {
	return fn(params)
}

// UserGetOneHandler interface for that can handle valid user get one params
type UserGetOneHandler interface {
	Handle(UserGetOneParams) middleware.Responder
}

// NewUserGetOne creates a new http.Handler for the user get one operation
func NewUserGetOne(ctx *middleware.Context, handler UserGetOneHandler) *UserGetOne {
	return &UserGetOne{Context: ctx, Handler: handler}
}

/*UserGetOne swagger:route GET /user/{nickname}/profile user userGetOne

Получение информации о пользователе

Получение информации о пользователе форума по его имени.


*/
type UserGetOne struct {
	Context *middleware.Context
	Handler UserGetOneHandler
}

func (o *UserGetOne) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUserGetOneParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
