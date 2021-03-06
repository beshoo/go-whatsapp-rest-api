// Code generated by go-swagger; DO NOT EDIT.

package hooks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PostMessageImageHandlerFunc turns a function with the right signature into a post message image handler
type PostMessageImageHandlerFunc func(PostMessageImageParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostMessageImageHandlerFunc) Handle(params PostMessageImageParams) middleware.Responder {
	return fn(params)
}

// PostMessageImageHandler interface for that can handle valid post message image params
type PostMessageImageHandler interface {
	Handle(PostMessageImageParams) middleware.Responder
}

// NewPostMessageImage creates a new http.Handler for the post message image operation
func NewPostMessageImage(ctx *middleware.Context, handler PostMessageImageHandler) *PostMessageImage {
	return &PostMessageImage{Context: ctx, Handler: handler}
}

/*PostMessageImage swagger:route POST /message/image Hooks postMessageImage

Image message hook

*/
type PostMessageImage struct {
	Context *middleware.Context
	Handler PostMessageImageHandler
}

func (o *PostMessageImage) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostMessageImageParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
