// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	models "AForum/app/generated/models"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
)

// NewUserCreateParams creates a new UserCreateParams object
// no default values defined in spec.
func NewUserCreateParams() UserCreateParams {

	return UserCreateParams{}
}

// UserCreateParams contains all the bound params for the user create operation
// typically these are obtained from a http.Request
//
// swagger:parameters userCreate
type UserCreateParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Идентификатор пользователя.
	  Required: true
	  In: path
	*/
	Nickname string
	/*Данные пользовательского профиля.
	  Required: true
	  In: body
	*/
	Profile *models.User
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewUserCreateParams() beforehand.
func (o *UserCreateParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rNickname, rhkNickname, _ := route.Params.GetOK("nickname")
	if err := o.bindNickname(rNickname, rhkNickname, route.Formats); err != nil {
		res = append(res, err)
	}

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.User
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("profile", "body"))
			} else {
				res = append(res, errors.NewParseError("profile", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Profile = &body
			}
		}
	} else {
		res = append(res, errors.Required("profile", "body"))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindNickname binds and validates parameter Nickname from path.
func (o *UserCreateParams) bindNickname(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.Nickname = raw

	return nil
}
