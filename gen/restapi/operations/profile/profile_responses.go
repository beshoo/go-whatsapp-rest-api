// Code generated by go-swagger; DO NOT EDIT.

package profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"bitbucket.org/rockyOO7/wa-api/gen/models"
)

// ProfileOKCode is the HTTP code returned for type ProfileOK
const ProfileOKCode int = 200

/*ProfileOK Profile Data

swagger:response profileOK
*/
type ProfileOK struct {

	/*
	  In: Body
	*/
	Payload *models.Profile `json:"body,omitempty"`
}

// NewProfileOK creates ProfileOK with default headers values
func NewProfileOK() *ProfileOK {

	return &ProfileOK{}
}

// WithPayload adds the payload to the profile o k response
func (o *ProfileOK) WithPayload(payload *models.Profile) *ProfileOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the profile o k response
func (o *ProfileOK) SetPayload(payload *models.Profile) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ProfileOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*ProfileDefault Error Response

swagger:response profileDefault
*/
type ProfileDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewProfileDefault creates ProfileDefault with default headers values
func NewProfileDefault(code int) *ProfileDefault {
	if code <= 0 {
		code = 500
	}

	return &ProfileDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the profile default response
func (o *ProfileDefault) WithStatusCode(code int) *ProfileDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the profile default response
func (o *ProfileDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the profile default response
func (o *ProfileDefault) WithPayload(payload *models.Error) *ProfileDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the profile default response
func (o *ProfileDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ProfileDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
