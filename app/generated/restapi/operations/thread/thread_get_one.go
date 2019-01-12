// Code generated by go-swagger; DO NOT EDIT.

package thread

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// ThreadGetOneHandlerFunc turns a function with the right signature into a thread get one handler
type ThreadGetOneHandlerFunc func(ThreadGetOneParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ThreadGetOneHandlerFunc) Handle(params ThreadGetOneParams) middleware.Responder {
	return fn(params)
}

// ThreadGetOneHandler interface for that can handle valid thread get one params
type ThreadGetOneHandler interface {
	Handle(ThreadGetOneParams) middleware.Responder
}

// NewThreadGetOne creates a new http.Handler for the thread get one operation
func NewThreadGetOne(ctx *middleware.Context, handler ThreadGetOneHandler) *ThreadGetOne {
	return &ThreadGetOne{Context: ctx, Handler: handler}
}

/*ThreadGetOne swagger:route GET /thread/{slug_or_id}/details thread threadGetOne

Получение информации о ветке обсуждения

Получение информации о ветке обсуждения по его имени.


*/
type ThreadGetOne struct {
	Context *middleware.Context
	Handler ThreadGetOneHandler
}

func (o *ThreadGetOne) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewThreadGetOneParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
