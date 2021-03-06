// Code generated by go-swagger; DO NOT EDIT.

package number

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"bitbucket.org/rockyOO7/wa-api/gen/models"
)

// GetAvatarOKCode is the HTTP code returned for type GetAvatarOK
const GetAvatarOKCode int = 200

/*GetAvatarOK Profile Data

swagger:response getAvatarOK
*/
type GetAvatarOK struct {

	/*
	  In: Body
	*/
	Payload *models.Profile `json:"body,omitempty"`
}

// NewGetAvatarOK creates GetAvatarOK with default headers values
func NewGetAvatarOK() *GetAvatarOK {

	return &GetAvatarOK{}
}

// WithPayload adds the payload to the get avatar o k response
func (o *GetAvatarOK) WithPayload(payload *models.Profile) *GetAvatarOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get avatar o k response
func (o *GetAvatarOK) SetPayload(payload *models.Profile) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAvatarOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetAvatarDefault Error Response

swagger:response getAvatarDefault
*/
type GetAvatarDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAvatarDefault creates GetAvatarDefault with default headers values
func NewGetAvatarDefault(code int) *GetAvatarDefault {
	if code <= 0 {
		code = 500
	}

	return &GetAvatarDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get avatar default response
func (o *GetAvatarDefault) WithStatusCode(code int) *GetAvatarDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get avatar default response
func (o *GetAvatarDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get avatar default response
func (o *GetAvatarDefault) WithPayload(payload *models.Error) *GetAvatarDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get avatar default response
func (o *GetAvatarDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAvatarDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
