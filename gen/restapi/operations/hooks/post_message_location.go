// Code generated by go-swagger; DO NOT EDIT.

package hooks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PostMessageLocationHandlerFunc turns a function with the right signature into a post message location handler
type PostMessageLocationHandlerFunc func(PostMessageLocationParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostMessageLocationHandlerFunc) Handle(params PostMessageLocationParams) middleware.Responder {
	return fn(params)
}

// PostMessageLocationHandler interface for that can handle valid post message location params
type PostMessageLocationHandler interface {
	Handle(PostMessageLocationParams) middleware.Responder
}

// NewPostMessageLocation creates a new http.Handler for the post message location operation
func NewPostMessageLocation(ctx *middleware.Context, handler PostMessageLocationHandler) *PostMessageLocation {
	return &PostMessageLocation{Context: ctx, Handler: handler}
}

/*PostMessageLocation swagger:route POST /message/location Hooks postMessageLocation

Location message hook

*/
type PostMessageLocation struct {
	Context *middleware.Context
	Handler PostMessageLocationHandler
}

func (o *PostMessageLocation) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostMessageLocationParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
