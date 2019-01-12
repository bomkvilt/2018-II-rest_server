// Code generated by go-swagger; DO NOT EDIT.

package thread

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// ThreadCreateHandlerFunc turns a function with the right signature into a thread create handler
type ThreadCreateHandlerFunc func(ThreadCreateParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ThreadCreateHandlerFunc) Handle(params ThreadCreateParams) middleware.Responder {
	return fn(params)
}

// ThreadCreateHandler interface for that can handle valid thread create params
type ThreadCreateHandler interface {
	Handle(ThreadCreateParams) middleware.Responder
}

// NewThreadCreate creates a new http.Handler for the thread create operation
func NewThreadCreate(ctx *middleware.Context, handler ThreadCreateHandler) *ThreadCreate {
	return &ThreadCreate{Context: ctx, Handler: handler}
}

/*ThreadCreate swagger:route POST /forum/{slug}/create thread threadCreate

Создание ветки

Добавление новой ветки обсуждения на форум.


*/
type ThreadCreate struct {
	Context *middleware.Context
	Handler ThreadCreateHandler
}

func (o *ThreadCreate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewThreadCreateParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
