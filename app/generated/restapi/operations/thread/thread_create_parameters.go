// Code generated by go-swagger; DO NOT EDIT.

package thread

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	models "Forum/app/generated/models"
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
)

// NewThreadCreateParams creates a new ThreadCreateParams object
// no default values defined in spec.
func NewThreadCreateParams() ThreadCreateParams {

	return ThreadCreateParams{}
}

// ThreadCreateParams contains all the bound params for the thread create operation
// typically these are obtained from a http.Request
//
// swagger:parameters threadCreate
type ThreadCreateParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Идентификатор форума.
	  Required: true
	  In: path
	*/
	Slug string
	/*Данные ветки обсуждения.
	  Required: true
	  In: body
	*/
	Thread *models.Thread
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewThreadCreateParams() beforehand.
func (o *ThreadCreateParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rSlug, rhkSlug, _ := route.Params.GetOK("slug")
	if err := o.bindSlug(rSlug, rhkSlug, route.Formats); err != nil {
		res = append(res, err)
	}

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.Thread
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("thread", "body"))
			} else {
				res = append(res, errors.NewParseError("thread", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Thread = &body
			}
		}
	} else {
		res = append(res, errors.Required("thread", "body"))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindSlug binds and validates parameter Slug from path.
func (o *ThreadCreateParams) bindSlug(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.Slug = raw

	return nil
}
