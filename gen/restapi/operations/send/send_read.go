// Code generated by go-swagger; DO NOT EDIT.

package send

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// SendReadHandlerFunc turns a function with the right signature into a send read handler
type SendReadHandlerFunc func(SendReadParams) middleware.Responder

// Handle executing the request and returning a response
func (fn SendReadHandlerFunc) Handle(params SendReadParams) middleware.Responder {
	return fn(params)
}

// SendReadHandler interface for that can handle valid send read params
type SendReadHandler interface {
	Handle(SendReadParams) middleware.Responder
}

// NewSendRead creates a new http.Handler for the send read operation
func NewSendRead(ctx *middleware.Context, handler SendReadHandler) *SendRead {
	return &SendRead{Context: ctx, Handler: handler}
}

/*SendRead swagger:route POST /send/ack/read Send sendRead

Send Read Reciept

*/
type SendRead struct {
	Context *middleware.Context
	Handler SendReadHandler
}

func (o *SendRead) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewSendReadParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
