// Code generated by go-swagger; DO NOT EDIT.

package hooks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PostNotifyLogoutHandlerFunc turns a function with the right signature into a post notify logout handler
type PostNotifyLogoutHandlerFunc func(PostNotifyLogoutParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostNotifyLogoutHandlerFunc) Handle(params PostNotifyLogoutParams) middleware.Responder {
	return fn(params)
}

// PostNotifyLogoutHandler interface for that can handle valid post notify logout params
type PostNotifyLogoutHandler interface {
	Handle(PostNotifyLogoutParams) middleware.Responder
}

// NewPostNotifyLogout creates a new http.Handler for the post notify logout operation
func NewPostNotifyLogout(ctx *middleware.Context, handler PostNotifyLogoutHandler) *PostNotifyLogout {
	return &PostNotifyLogout{Context: ctx, Handler: handler}
}

/*PostNotifyLogout swagger:route POST /notify/logout Hooks postNotifyLogout

Notify when user logs out

*/
type PostNotifyLogout struct {
	Context *middleware.Context
	Handler PostNotifyLogoutHandler
}

func (o *PostNotifyLogout) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostNotifyLogoutParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
