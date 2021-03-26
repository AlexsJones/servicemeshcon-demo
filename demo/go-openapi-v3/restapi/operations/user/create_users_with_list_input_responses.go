// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// CreateUsersWithListInputBadRequestCode is the HTTP code returned for type CreateUsersWithListInputBadRequest
const CreateUsersWithListInputBadRequestCode int = 400

/*CreateUsersWithListInputBadRequest Error creating user

swagger:response createUsersWithListInputBadRequest
*/
type CreateUsersWithListInputBadRequest struct {
}

// NewCreateUsersWithListInputBadRequest creates CreateUsersWithListInputBadRequest with default headers values
func NewCreateUsersWithListInputBadRequest() *CreateUsersWithListInputBadRequest {

	return &CreateUsersWithListInputBadRequest{}
}

// WriteResponse to the client
func (o *CreateUsersWithListInputBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

/*CreateUsersWithListInputDefault successful operation

swagger:response createUsersWithListInputDefault
*/
type CreateUsersWithListInputDefault struct {
	_statusCode int
}

// NewCreateUsersWithListInputDefault creates CreateUsersWithListInputDefault with default headers values
func NewCreateUsersWithListInputDefault(code int) *CreateUsersWithListInputDefault {
	if code <= 0 {
		code = 500
	}

	return &CreateUsersWithListInputDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create users with list input default response
func (o *CreateUsersWithListInputDefault) WithStatusCode(code int) *CreateUsersWithListInputDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create users with list input default response
func (o *CreateUsersWithListInputDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WriteResponse to the client
func (o *CreateUsersWithListInputDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(o._statusCode)
}
