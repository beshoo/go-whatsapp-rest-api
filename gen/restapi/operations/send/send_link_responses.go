// Code generated by go-swagger; DO NOT EDIT.

package send

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"bitbucket.org/rockyOO7/wa-api/gen/models"
)

// SendLinkOKCode is the HTTP code returned for type SendLinkOK
const SendLinkOKCode int = 200

/*SendLinkOK Returns status of broadcast

swagger:response sendLinkOK
*/
type SendLinkOK struct {

	/*
	  In: Body
	*/
	Payload *models.BroadcastStatus `json:"body,omitempty"`
}

// NewSendLinkOK creates SendLinkOK with default headers values
func NewSendLinkOK() *SendLinkOK {

	return &SendLinkOK{}
}

// WithPayload adds the payload to the send link o k response
func (o *SendLinkOK) WithPayload(payload *models.BroadcastStatus) *SendLinkOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the send link o k response
func (o *SendLinkOK) SetPayload(payload *models.BroadcastStatus) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SendLinkOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*SendLinkDefault Error Response

swagger:response sendLinkDefault
*/
type SendLinkDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewSendLinkDefault creates SendLinkDefault with default headers values
func NewSendLinkDefault(code int) *SendLinkDefault {
	if code <= 0 {
		code = 500
	}

	return &SendLinkDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the send link default response
func (o *SendLinkDefault) WithStatusCode(code int) *SendLinkDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the send link default response
func (o *SendLinkDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the send link default response
func (o *SendLinkDefault) WithPayload(payload *models.Error) *SendLinkDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the send link default response
func (o *SendLinkDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SendLinkDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}