// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	models "AForum/app/generated/models"
	"github.com/go-openapi/runtime"
)

// UserGetOneOKCode is the HTTP code returned for type UserGetOneOK
const UserGetOneOKCode int = 200

/*UserGetOneOK Информация о пользователе.


swagger:response userGetOneOK
*/
type UserGetOneOK struct {

	/*
	  In: Body
	*/
	Payload *models.User `json:"body,omitempty"`
}

// NewUserGetOneOK creates UserGetOneOK with default headers values
func NewUserGetOneOK() *UserGetOneOK {

	return &UserGetOneOK{}
}

// WithPayload adds the payload to the user get one o k response
func (o *UserGetOneOK) WithPayload(payload *models.User) *UserGetOneOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the user get one o k response
func (o *UserGetOneOK) SetPayload(payload *models.User) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UserGetOneOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UserGetOneNotFoundCode is the HTTP code returned for type UserGetOneNotFound
const UserGetOneNotFoundCode int = 404

/*UserGetOneNotFound Пользователь отсутсвует в системе.


swagger:response userGetOneNotFound
*/
type UserGetOneNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUserGetOneNotFound creates UserGetOneNotFound with default headers values
func NewUserGetOneNotFound() *UserGetOneNotFound {

	return &UserGetOneNotFound{}
}

// WithPayload adds the payload to the user get one not found response
func (o *UserGetOneNotFound) WithPayload(payload *models.Error) *UserGetOneNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the user get one not found response
func (o *UserGetOneNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UserGetOneNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
