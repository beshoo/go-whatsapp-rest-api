// Code generated by go-swagger; DO NOT EDIT.

package number

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"bitbucket.org/rockyOO7/wa-api/gen/models"
)

// IsOnlineOKCode is the HTTP code returned for type IsOnlineOK
const IsOnlineOKCode int = 200

/*IsOnlineOK Return if the user is online

swagger:response isOnlineOK
*/
type IsOnlineOK struct {

	/*
	  In: Body
	*/
	Payload *IsOnlineOKBody `json:"body,omitempty"`
}

// NewIsOnlineOK creates IsOnlineOK with default headers values
func NewIsOnlineOK() *IsOnlineOK {

	return &IsOnlineOK{}
}

// WithPayload adds the payload to the is online o k response
func (o *IsOnlineOK) WithPayload(payload *IsOnlineOKBody) *IsOnlineOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the is online o k response
func (o *IsOnlineOK) SetPayload(payload *IsOnlineOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *IsOnlineOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*IsOnlineDefault Error Response

swagger:response isOnlineDefault
*/
type IsOnlineDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewIsOnlineDefault creates IsOnlineDefault with default headers values
func NewIsOnlineDefault(code int) *IsOnlineDefault {
	if code <= 0 {
		code = 500
	}

	return &IsOnlineDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the is online default response
func (o *IsOnlineDefault) WithStatusCode(code int) *IsOnlineDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the is online default response
func (o *IsOnlineDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the is online default response
func (o *IsOnlineDefault) WithPayload(payload *models.Error) *IsOnlineDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the is online default response
func (o *IsOnlineDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *IsOnlineDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
