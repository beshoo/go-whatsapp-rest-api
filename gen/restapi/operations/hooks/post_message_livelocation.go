// Code generated by go-swagger; DO NOT EDIT.

package hooks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PostMessageLivelocationHandlerFunc turns a function with the right signature into a post message livelocation handler
type PostMessageLivelocationHandlerFunc func(PostMessageLivelocationParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostMessageLivelocationHandlerFunc) Handle(params PostMessageLivelocationParams) middleware.Responder {
	return fn(params)
}

// PostMessageLivelocationHandler interface for that can handle valid post message livelocation params
type PostMessageLivelocationHandler interface {
	Handle(PostMessageLivelocationParams) middleware.Responder
}

// NewPostMessageLivelocation creates a new http.Handler for the post message livelocation operation
func NewPostMessageLivelocation(ctx *middleware.Context, handler PostMessageLivelocationHandler) *PostMessageLivelocation {
	return &PostMessageLivelocation{Context: ctx, Handler: handler}
}

/*PostMessageLivelocation swagger:route POST /message/livelocation Hooks postMessageLivelocation

Live Location message hook

*/
type PostMessageLivelocation struct {
	Context *middleware.Context
	Handler PostMessageLivelocationHandler
}

func (o *PostMessageLivelocation) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostMessageLivelocationParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}