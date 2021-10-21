// Code generated by go-swagger; DO NOT EDIT.

package send

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"bitbucket.org/rockyOO7/wa-api/gen/models"
)

// SendAudioOKCode is the HTTP code returned for type SendAudioOK
const SendAudioOKCode int = 200

/*SendAudioOK Returns status of broadcast

swagger:response sendAudioOK
*/
type SendAudioOK struct {

	/*
	  In: Body
	*/
	Payload *models.BroadcastStatus `json:"body,omitempty"`
}

// NewSendAudioOK creates SendAudioOK with default headers values
func NewSendAudioOK() *SendAudioOK {

	return &SendAudioOK{}
}

// WithPayload adds the payload to the send audio o k response
func (o *SendAudioOK) WithPayload(payload *models.BroadcastStatus) *SendAudioOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the send audio o k response
func (o *SendAudioOK) SetPayload(payload *models.BroadcastStatus) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SendAudioOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*SendAudioDefault Error Response

swagger:response sendAudioDefault
*/
type SendAudioDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewSendAudioDefault creates SendAudioDefault with default headers values
func NewSendAudioDefault(code int) *SendAudioDefault {
	if code <= 0 {
		code = 500
	}

	return &SendAudioDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the send audio default response
func (o *SendAudioDefault) WithStatusCode(code int) *SendAudioDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the send audio default response
func (o *SendAudioDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the send audio default response
func (o *SendAudioDefault) WithPayload(payload *models.Error) *SendAudioDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the send audio default response
func (o *SendAudioDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SendAudioDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}