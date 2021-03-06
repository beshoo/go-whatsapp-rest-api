// Code generated by go-swagger; DO NOT EDIT.

package send

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"bitbucket.org/rockyOO7/wa-api/gen/models"
)

// SendAudioRecordOKCode is the HTTP code returned for type SendAudioRecordOK
const SendAudioRecordOKCode int = 200

/*SendAudioRecordOK Returns status of broadcast

swagger:response sendAudioRecordOK
*/
type SendAudioRecordOK struct {

	/*
	  In: Body
	*/
	Payload *models.BroadcastStatus `json:"body,omitempty"`
}

// NewSendAudioRecordOK creates SendAudioRecordOK with default headers values
func NewSendAudioRecordOK() *SendAudioRecordOK {

	return &SendAudioRecordOK{}
}

// WithPayload adds the payload to the send audio record o k response
func (o *SendAudioRecordOK) WithPayload(payload *models.BroadcastStatus) *SendAudioRecordOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the send audio record o k response
func (o *SendAudioRecordOK) SetPayload(payload *models.BroadcastStatus) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SendAudioRecordOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*SendAudioRecordDefault Error Response

swagger:response sendAudioRecordDefault
*/
type SendAudioRecordDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewSendAudioRecordDefault creates SendAudioRecordDefault with default headers values
func NewSendAudioRecordDefault(code int) *SendAudioRecordDefault {
	if code <= 0 {
		code = 500
	}

	return &SendAudioRecordDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the send audio record default response
func (o *SendAudioRecordDefault) WithStatusCode(code int) *SendAudioRecordDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the send audio record default response
func (o *SendAudioRecordDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the send audio record default response
func (o *SendAudioRecordDefault) WithPayload(payload *models.Error) *SendAudioRecordDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the send audio record default response
func (o *SendAudioRecordDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SendAudioRecordDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
