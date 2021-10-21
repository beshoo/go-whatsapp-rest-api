// Code generated by go-swagger; DO NOT EDIT.

package send

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// SendImageHandlerFunc turns a function with the right signature into a send image handler
type SendImageHandlerFunc func(SendImageParams) middleware.Responder

// Handle executing the request and returning a response
func (fn SendImageHandlerFunc) Handle(params SendImageParams) middleware.Responder {
	return fn(params)
}

// SendImageHandler interface for that can handle valid send image params
type SendImageHandler interface {
	Handle(SendImageParams) middleware.Responder
}

// NewSendImage creates a new http.Handler for the send image operation
func NewSendImage(ctx *middleware.Context, handler SendImageHandler) *SendImage {
	return &SendImage{Context: ctx, Handler: handler}
}

/*SendImage swagger:route POST /send/image Send sendImage

Send Image Message

*/
type SendImage struct {
	Context *middleware.Context
	Handler SendImageHandler
}

func (o *SendImage) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewSendImageParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}