// Code generated by go-swagger; DO NOT EDIT.

package profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetContactsHandlerFunc turns a function with the right signature into a get contacts handler
type GetContactsHandlerFunc func(GetContactsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetContactsHandlerFunc) Handle(params GetContactsParams) middleware.Responder {
	return fn(params)
}

// GetContactsHandler interface for that can handle valid get contacts params
type GetContactsHandler interface {
	Handle(GetContactsParams) middleware.Responder
}

// NewGetContacts creates a new http.Handler for the get contacts operation
func NewGetContacts(ctx *middleware.Context, handler GetContactsHandler) *GetContacts {
	return &GetContacts{Context: ctx, Handler: handler}
}

/*GetContacts swagger:route GET /profile/contacts Profile getContacts

Get Contacts for the user

*/
type GetContacts struct {
	Context *middleware.Context
	Handler GetContactsHandler
}

func (o *GetContacts) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetContactsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
