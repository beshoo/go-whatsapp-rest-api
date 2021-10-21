// Code generated by go-swagger; DO NOT EDIT.

package profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DisconnectHandlerFunc turns a function with the right signature into a disconnect handler
type DisconnectHandlerFunc func(DisconnectParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DisconnectHandlerFunc) Handle(params DisconnectParams) middleware.Responder {
	return fn(params)
}

// DisconnectHandler interface for that can handle valid disconnect params
type DisconnectHandler interface {
	Handle(DisconnectParams) middleware.Responder
}

// NewDisconnect creates a new http.Handler for the disconnect operation
func NewDisconnect(ctx *middleware.Context, handler DisconnectHandler) *Disconnect {
	return &Disconnect{Context: ctx, Handler: handler}
}

/*Disconnect swagger:route POST /profile/phone/disconnect Profile disconnect

Disconnect Session Id used while succesfull scan

*/
type Disconnect struct {
	Context *middleware.Context
	Handler DisconnectHandler
}

func (o *Disconnect) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDisconnectParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
