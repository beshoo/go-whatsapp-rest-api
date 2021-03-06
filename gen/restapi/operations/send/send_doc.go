// Code generated by go-swagger; DO NOT EDIT.

package send

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// SendDocHandlerFunc turns a function with the right signature into a send doc handler
type SendDocHandlerFunc func(SendDocParams) middleware.Responder

// Handle executing the request and returning a response
func (fn SendDocHandlerFunc) Handle(params SendDocParams) middleware.Responder {
	return fn(params)
}

// SendDocHandler interface for that can handle valid send doc params
type SendDocHandler interface {
	Handle(SendDocParams) middleware.Responder
}

// NewSendDoc creates a new http.Handler for the send doc operation
func NewSendDoc(ctx *middleware.Context, handler SendDocHandler) *SendDoc {
	return &SendDoc{Context: ctx, Handler: handler}
}

/*SendDoc swagger:route POST /send/doc Send sendDoc

Send doc message

*/
type SendDoc struct {
	Context *middleware.Context
	Handler SendDocHandler
}

func (o *SendDoc) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewSendDocParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
