// Code generated by go-swagger; DO NOT EDIT.

package post

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	models "github.com/bomkvilt/tech-db-app/app/generated/models"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewPostUpdateParams creates a new PostUpdateParams object
// no default values defined in spec.
func NewPostUpdateParams() PostUpdateParams {

	return PostUpdateParams{}
}

// PostUpdateParams contains all the bound params for the post update operation
// typically these are obtained from a http.Request
//
// swagger:parameters postUpdate
type PostUpdateParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Идентификатор сообщения.
	  Required: true
	  In: path
	*/
	ID int64
	/*Изменения сообщения.
	  Required: true
	  In: body
	*/
	Post *models.PostUpdate
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPostUpdateParams() beforehand.
func (o *PostUpdateParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rID, rhkID, _ := route.Params.GetOK("id")
	if err := o.bindID(rID, rhkID, route.Formats); err != nil {
		res = append(res, err)
	}

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.PostUpdate
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("post", "body"))
			} else {
				res = append(res, errors.NewParseError("post", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Post = &body
			}
		}
	} else {
		res = append(res, errors.Required("post", "body"))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindID binds and validates parameter ID from path.
func (o *PostUpdateParams) bindID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("id", "path", "int64", raw)
	}
	o.ID = value

	return nil
}