// Code generated by go-swagger; DO NOT EDIT.

package send

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"bitbucket.org/rockyOO7/wa-api/gen/models"
)

// SendImageOKCode is the HTTP code returned for type SendImageOK
const SendImageOKCode int = 200

/*SendImageOK Returns status of broadcast

swagger:response sendImageOK
*/
type SendImageOK struct {

	/*
	  In: Body
	*/
	Payload *models.BroadcastStatus `json:"body,omitempty"`
}

// NewSendImageOK creates SendImageOK with default headers values
func NewSendImageOK() *SendImageOK {

	return &SendImageOK{}
}

// WithPayload adds the payload to the send image o k response
func (o *SendImageOK) WithPayload(payload *models.BroadcastStatus) *SendImageOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the send image o k response
func (o *SendImageOK) SetPayload(payload *models.BroadcastStatus) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SendImageOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*SendImageDefault Error Response

swagger:response sendImageDefault
*/
type SendImageDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewSendImageDefault creates SendImageDefault with default headers values
func NewSendImageDefault(code int) *SendImageDefault {
	if code <= 0 {
		code = 500
	}

	return &SendImageDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the send image default response
func (o *SendImageDefault) WithStatusCode(code int) *SendImageDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the send image default response
func (o *SendImageDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the send image default response
func (o *SendImageDefault) WithPayload(payload *models.Error) *SendImageDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the send image default response
func (o *SendImageDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SendImageDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}