// Code generated by go-swagger; DO NOT EDIT.

package send

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// SendAudioHandlerFunc turns a function with the right signature into a send audio handler
type SendAudioHandlerFunc func(SendAudioParams) middleware.Responder

// Handle executing the request and returning a response
func (fn SendAudioHandlerFunc) Handle(params SendAudioParams) middleware.Responder {
	return fn(params)
}

// SendAudioHandler interface for that can handle valid send audio params
type SendAudioHandler interface {
	Handle(SendAudioParams) middleware.Responder
}

// NewSendAudio creates a new http.Handler for the send audio operation
func NewSendAudio(ctx *middleware.Context, handler SendAudioHandler) *SendAudio {
	return &SendAudio{Context: ctx, Handler: handler}
}

/*SendAudio swagger:route POST /send/audio Send sendAudio

Send audio message

*/
type SendAudio struct {
	Context *middleware.Context
	Handler SendAudioHandler
}

func (o *SendAudio) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewSendAudioParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}